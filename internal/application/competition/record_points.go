package competition

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type RecordPointsCommand struct {
	UserID    string
	Minutes   int
	Timestamp time.Time
}

type RecordPointsHandler struct {
	Repo ports.CompetitionRepository
}

func NewRecordPointsHandler(repo ports.CompetitionRepository) *RecordPointsHandler {
	return &RecordPointsHandler{Repo: repo}
}

// This function will be called from LogReadingHandler.
// It finds all active competitions and awards points for the reading session.
func (h *RecordPointsHandler) Handle(cmd RecordPointsCommand) error {

	comps, err := h.Repo.FindActive(cmd.Timestamp)
	if err != nil {
		return core.New(core.ServerError, "failed to load competitions")
	}

	for _, cmp := range comps {
		participant, exists := cmp.Participants[cmd.UserID]
		if !exists {
			continue // user is not part of this competition
		}

		participant.AddReading(cmd.Minutes, cmd.Timestamp, cmp.Rules)

		if err := h.Repo.Save(cmp); err != nil {
			return core.New(core.ServerError, "failed to update competition")
		}
	}

	return nil
}
