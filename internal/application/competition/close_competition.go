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

func (h *CloseCompetitionHandler) Handle(cmd CloseCompetitionCommand) ([]string, error) {

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

	// Sort participants by points
	type scored struct {
		UserID string
		Points int
	}

	list := make([]scored, 0, len(cmp.Participants))
	for _, p := range cmp.Participants {
		list = append(list, scored{UserID: p.UserID, Points: p.Points})
	}

	sort.Slice(list, func(a, b int) bool {
		return list[a].Points > list[b].Points
	})

	// Collect winners in sorted order
	winners := make([]string, len(list))
	for i, s := range list {
		winners[i] = s.UserID
	}

	if err := h.Repo.Save(cmp); err != nil {
		return nil, core.New(core.ServerError, "failed to save competition")
	}

	return winners, nil
}
