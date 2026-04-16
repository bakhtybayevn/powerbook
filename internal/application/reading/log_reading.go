package reading

import (
	"context"
	"fmt"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/reading"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type LogReadingCommand struct {
	UserID    string
	Minutes   int
	Source    string
	Timestamp time.Time
}

type LogReadingHandler struct {
	UserRepo        ports.UserRepository
	ReadingRepo     ports.ReadingRepository
	CompetitionRepo ports.CompetitionRepository
	Leaderboard     ports.LeaderboardPort // NEW
}

func NewLogReadingHandler(
	userRepo ports.UserRepository,
	readingRepo ports.ReadingRepository,
	competitionRepo ports.CompetitionRepository,
	leaderboard ports.LeaderboardPort,
) *LogReadingHandler {
	return &LogReadingHandler{
		UserRepo:        userRepo,
		ReadingRepo:     readingRepo,
		CompetitionRepo: competitionRepo,
		Leaderboard:     leaderboard,
	}
}

func (h *LogReadingHandler) Handle(cmd LogReadingCommand) (newStreak int, totalMinutes int, err error) {
	// validation
	if cmd.Minutes <= 0 {
		return 0, 0, core.New(core.ValidationError, "minutes must be > 0")
	}
	if cmd.Source == "" {
		cmd.Source = "unknown"
	}
	if cmd.Minutes > 1440 {
		return 0, 0, core.New(core.ValidationError, "minutes cannot exceed 1440 (24 hours)")
	}

	now := time.Now().UTC()
	if cmd.Timestamp.After(now) {
		return 0, 0, core.New(core.ValidationError, "timestamp cannot be in the future")
	}

	// Check daily cap: max 1440 minutes per day
	dayStart := time.Date(cmd.Timestamp.Year(), cmd.Timestamp.Month(), cmd.Timestamp.Day(), 0, 0, 0, 0, time.UTC)
	dayEnd := dayStart.Add(24 * time.Hour)
	dayLogs, _ := h.ReadingRepo.ListByDateRange(cmd.UserID, dayStart, dayEnd)
	dayTotal := 0
	for _, dl := range dayLogs {
		dayTotal += dl.Minutes
	}
	if dayTotal+cmd.Minutes > 1440 {
		remaining := 1440 - dayTotal
		if remaining <= 0 {
			return 0, 0, core.New(core.ValidationError, "daily reading limit reached (1440 min/day)")
		}
		return 0, 0, core.New(core.ValidationError, fmt.Sprintf("would exceed daily limit; you can log %d more minutes today", remaining))
	}

	// load user
	u, err := h.UserRepo.Get(cmd.UserID)
	if err != nil {
		return 0, 0, core.New(core.NotFoundError, "user not found")
	}

	// domain logic - update user streak
	newStreak, totalMinutes = u.LogReading(cmd.Minutes, cmd.Timestamp)

	// persist reading log
	rd := reading.NewReading(cmd.UserID, cmd.Minutes, cmd.Source, cmd.Timestamp.UTC())
	if err := h.ReadingRepo.Save(rd); err != nil {
		return 0, 0, core.New(core.ServerError, "failed to save reading")
	}

	// save updated user
	if err := h.UserRepo.Save(u); err != nil {
		return 0, 0, core.New(core.ServerError, "failed to update user")
	}

	// === AWARD POINTS TO COMPETITIONS ===
	activeComps, err := h.CompetitionRepo.FindActive(cmd.Timestamp)
	if err == nil {
		for _, cmp := range activeComps {
			participant, ok := cmp.Participants[cmd.UserID]
			if !ok {
				continue
			}

			// update in competition object
			participant.AddReading(cmd.Minutes, cmd.Timestamp, cmp.Rules)
			_ = h.CompetitionRepo.Save(cmp)

			// compute points
			points := float64(cmp.Rules.PointsPerMinute * cmd.Minutes)

			// push to Redis leaderboard (best effort)
			_, _ = h.Leaderboard.AddScore(context.Background(), cmp.ID, cmd.UserID, points)
		}
	}

	return newStreak, totalMinutes, nil
}
