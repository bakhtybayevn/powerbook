package reading

import (
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
}

func NewLogReadingHandler(userRepo ports.UserRepository, readingRepo ports.ReadingRepository, competitionRepo ports.CompetitionRepository) *LogReadingHandler {
	return &LogReadingHandler{
		UserRepo:        userRepo,
		ReadingRepo:     readingRepo,
		CompetitionRepo: competitionRepo,
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

	// timestamp cannot be in the future
	now := time.Now().UTC()
	if cmd.Timestamp.After(now) {
		return 0, 0, core.New(core.ValidationError, "timestamp cannot be in the future")
	}

	// load user
	u, err := h.UserRepo.Get(cmd.UserID)
	if err != nil {
		return 0, 0, core.New(core.NotFoundError, "user not found")
	}

	// apply domain logic on user
	newStreak, totalMinutes = u.LogReading(cmd.Minutes, cmd.Timestamp)

	// persist reading entry
	rd := &reading.Reading{
		UserID:    cmd.UserID,
		Minutes:   cmd.Minutes,
		Source:    cmd.Source,
		Timestamp: cmd.Timestamp.UTC(),
	}
	if err := h.ReadingRepo.Save(rd); err != nil {
		return 0, 0, core.New(core.ServerError, "failed to save reading")
	}

	// persist updated user (save streak/total)
	if err := h.UserRepo.Save(u); err != nil {
		// if we failed to save user, it's a server error
		return 0, 0, core.New(core.ServerError, "failed to update user")
	}

	activeComps, err := h.CompetitionRepo.FindActive(cmd.Timestamp)
	if err == nil {
		for _, cmp := range activeComps {
			participant, ok := cmp.Participants[cmd.UserID]
			if !ok {
				continue
			}

			participant.AddReading(cmd.Minutes, cmd.Timestamp, cmp.Rules)
			_ = h.CompetitionRepo.Save(cmp)
		}
	}

	// optionally: publish domain events (ReadingLogged, StreakUpdated) — omitted for in-memory version
	return newStreak, totalMinutes, nil
}
