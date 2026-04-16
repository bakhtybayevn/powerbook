package competition

import (
	"math/rand"
	"sort"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type CloseCompetitionCommand struct {
	CompetitionID string
}

type CloseCompetitionHandler struct {
	Repo     ports.CompetitionRepository
	UserRepo ports.UserRepository
}

func NewCloseCompetitionHandler(repo ports.CompetitionRepository, userRepo ports.UserRepository) *CloseCompetitionHandler {
	return &CloseCompetitionHandler{Repo: repo, UserRepo: userRepo}
}

type Winner struct {
	UserID       string `json:"user_id"`
	Points       int    `json:"points"`
	DaysRead     int    `json:"days_read"`
	MinutesTotal int    `json:"minutes_total"`
	Rank         int    `json:"rank"`
	XPEarned     int    `json:"xp_earned"`
}

// XP award logic:
// - 1st place: 200 XP
// - 2nd place: 150 XP
// - 3rd place: 100 XP
// - Top 50%: 50 XP
// - Bottom 50%: 20 XP (participation)
// - Bonus: +5 XP per day read during competition
// - Streak bonus: +30 XP if read every day of competition
func calculateXP(rank, total, daysRead int, compDays int) int {
	xp := 0

	// Position-based XP
	switch rank {
	case 1:
		xp += 200
	case 2:
		xp += 150
	case 3:
		xp += 100
	default:
		mid := total / 2
		if total%2 != 0 {
			mid++
		}
		if rank <= mid {
			xp += 50 // top 50%
		} else {
			xp += 20 // bottom 50%
		}
	}

	// Per-day bonus
	xp += daysRead * 5

	// Perfect streak bonus (read every day of competition)
	if compDays > 0 && daysRead >= compDays {
		xp += 30
	}

	return xp
}

func (h *CloseCompetitionHandler) Handle(cmd CloseCompetitionCommand) ([]Winner, []*competition.GiftExchange, error) {
	if cmd.CompetitionID == "" {
		return nil, nil, core.New(core.ValidationError, "competition id is required")
	}

	cmp, err := h.Repo.Get(cmd.CompetitionID)
	if err != nil {
		return nil, nil, core.New(core.NotFoundError, "competition not found")
	}

	if cmp.Status == competition.StatusClosed {
		return nil, nil, core.New(core.ValidationError, "competition already closed")
	}

	// Mark closed
	cmp.Status = competition.StatusClosed

	// Build sorted standings
	winners := make([]Winner, 0, len(cmp.Participants))
	for _, p := range cmp.Participants {
		winners = append(winners, Winner{
			UserID:       p.UserID,
			Points:       p.Points,
			DaysRead:     p.DaysRead,
			MinutesTotal: p.MinutesTotal,
		})
	}

	sort.Slice(winners, func(i, j int) bool {
		return winners[i].Points > winners[j].Points
	})

	// Calculate competition duration in days
	compDays := int(cmp.EndDate.Sub(cmp.StartDate).Hours()/24) + 1

	// Assign ranks and XP
	total := len(winners)
	for i := range winners {
		winners[i].Rank = i + 1
		winners[i].XPEarned = calculateXP(i+1, total, winners[i].DaysRead, compDays)

		// Award XP to user
		u, err := h.UserRepo.Get(winners[i].UserID)
		if err == nil {
			u.AddXP(winners[i].XPEarned)
			h.UserRepo.Save(u)
		}
	}

	// Generate gift pairings: top 50% gives to bottom 50%
	var gifts []*competition.GiftExchange
	if total >= 2 {
		mid := total / 2
		if total%2 != 0 {
			mid = total/2 + 1
		}
		topHalf := make([]string, 0, mid)
		bottomHalf := make([]string, 0, total-mid)

		for i, w := range winners {
			if i < mid {
				topHalf = append(topHalf, w.UserID)
			} else {
				bottomHalf = append(bottomHalf, w.UserID)
			}
		}

		// Shuffle bottom half for random pairing
		rand.Shuffle(len(bottomHalf), func(i, j int) {
			bottomHalf[i], bottomHalf[j] = bottomHalf[j], bottomHalf[i]
		})

		// Pair: each top gives to one bottom (limited by smaller group)
		pairCount := len(bottomHalf)
		if len(topHalf) < pairCount {
			pairCount = len(topHalf)
		}

		for i := 0; i < pairCount; i++ {
			g := competition.NewGiftExchange(cmp.ID, topHalf[i], bottomHalf[i])
			if err := h.Repo.SaveGiftExchange(g); err == nil {
				gifts = append(gifts, g)
			}
		}
	}

	if err := h.Repo.Save(cmp); err != nil {
		return nil, nil, core.New(core.ServerError, "failed to save competition")
	}

	return winners, gifts, nil
}
