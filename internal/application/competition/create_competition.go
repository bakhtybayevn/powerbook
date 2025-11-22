package competition

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type CreateCompetitionCommand struct {
	Name            string
	StartDate       time.Time
	EndDate         time.Time
	PointsPerMinute int
}

type CreateCompetitionHandler struct {
	Repo ports.CompetitionRepository
}

func NewCreateCompetitionHandler(repo ports.CompetitionRepository) *CreateCompetitionHandler {
	return &CreateCompetitionHandler{Repo: repo}
}

func (h *CreateCompetitionHandler) Handle(cmd CreateCompetitionCommand) (*competition.Competition, error) {

	// === VALIDATION ===
	if cmd.Name == "" {
		return nil, core.New(core.ValidationError, "competition name is required")
	}

	if cmd.PointsPerMinute <= 0 {
		return nil, core.New(core.ValidationError, "points_per_minute must be > 0")
	}

	if cmd.EndDate.Before(cmd.StartDate) {
		return nil, core.New(core.ValidationError, "end_date cannot be before start_date")
	}

	rules := competition.Rules{
		PointsPerMinute: cmd.PointsPerMinute,
	}

	cmp, err := competition.NewCompetition(cmd.Name, cmd.StartDate, cmd.EndDate, rules)
	if err != nil {
		return nil, core.New(core.ValidationError, err.Error())
	}

	if err := h.Repo.Create(cmp); err != nil {
		return nil, core.New(core.ServerError, "failed to save competition")
	}

	return cmp, nil
}
