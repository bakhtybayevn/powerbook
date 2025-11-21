package competition

import (
	"sort"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type CloseCompetitionCommand struct {
	CompetitionID string
}

type CloseCompetitionHandler struct {
	Repo ports.CompetitionRepository
}

func NewCloseCompetitionHandler(repo ports.CompetitionRepository) *CloseCompetitionHandler {
	return &CloseCompetitionHandler{Repo: repo}
}

type Winner struct {
	UserID       string `json:"user_id"`
	Points       int    `json:"points"`
	DaysRead     int    `json:"days_read"`
	MinutesTotal int    `json:"minutes_total"`
}

func (h *CloseCompetitionHandler) Handle(cmd CloseCompetitionCommand) ([]Winner, error) {

	if cmd.CompetitionID == "" {
		return nil, core.New(core.ValidationError, "competition id is required")
	}

	cmp, err := h.Repo.Get(cmd.CompetitionID)
	if err != nil {
		return nil, core.New(core.NotFoundError, "competition not found")
	}

	if cmp.Status == competition.StatusClosed {
		return nil, core.New(core.ValidationError, "competition already closed")
	}

	// Mark closed
	cmp.Status = competition.StatusClosed

	// Build winners structure and sort
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

	if err := h.Repo.Save(cmp); err != nil {
		return nil, core.New(core.ServerError, "failed to save competition")
	}

	return winners, nil
}
