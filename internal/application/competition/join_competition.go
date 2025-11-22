package competition

import (
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type JoinCompetitionCommand struct {
	UserID        string
	CompetitionID string
}

type JoinCompetitionHandler struct {
	Repo ports.CompetitionRepository
}

func NewJoinCompetitionHandler(repo ports.CompetitionRepository) *JoinCompetitionHandler {
	return &JoinCompetitionHandler{Repo: repo}
}

func (h *JoinCompetitionHandler) Handle(cmd JoinCompetitionCommand) error {
	if cmd.CompetitionID == "" {
		return core.New(core.ValidationError, "competition id is required")
	}
	if cmd.UserID == "" {
		return core.New(core.ValidationError, "user id is required")
	}

	cmp, err := h.Repo.Get(cmd.CompetitionID)
	if err != nil {
		return core.New(core.NotFoundError, "competition not found")
	}

	if cmp.Status == competition.StatusClosed {
		return core.New(core.ValidationError, "competition is closed")
	}

	// Prevent duplicate join
	if _, exists := cmp.Participants[cmd.UserID]; exists {
		return core.New(core.ValidationError, "user already joined this competition")
	}

	// Add participant
	p := competition.NewParticipant(cmd.UserID)
    cmp.Participants[cmd.UserID] = p

	if err := h.Repo.SaveParticipant(cmd.CompetitionID, p); err != nil {
        return core.New(core.ServerError, "failed to add participant")
    }

	return nil
}
