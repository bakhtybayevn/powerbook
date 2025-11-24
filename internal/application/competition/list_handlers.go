package competition

import (
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/domain/user"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type ListAllCompetitionsHandler struct {
	CompRepo ports.CompetitionRepository
	UserRepo ports.UserRepository
}

func NewListAllCompetitionsHandler(comp ports.CompetitionRepository, users ports.UserRepository) *ListAllCompetitionsHandler {
	return &ListAllCompetitionsHandler{CompRepo: comp, UserRepo: users}
}

func (h *ListAllCompetitionsHandler) Handle() ([]*competition.Competition, map[string]*user.User, error) {
	comps, err := h.CompRepo.GetAll()
	if err != nil {
		return nil, nil, err
	}

	allUsers := map[string]*user.User{}
	for _, c := range comps {
		for uid := range c.Participants {
			if _, exists := allUsers[uid]; !exists {
				u, err := h.UserRepo.Get(uid)
				if err == nil {
					allUsers[uid] = u
				}
			}
		}
	}

	return comps, allUsers, nil
}

// -------------------------------------

type ListMyCompetitionsHandler struct {
	CompRepo ports.CompetitionRepository
	UserRepo ports.UserRepository
}

func NewListMyCompetitionsHandler(comp ports.CompetitionRepository, users ports.UserRepository) *ListMyCompetitionsHandler {
	return &ListMyCompetitionsHandler{CompRepo: comp, UserRepo: users}
}

type ListMyCommand struct {
	UserID string
}

func (h *ListMyCompetitionsHandler) Handle(cmd ListMyCommand) ([]*competition.Competition, map[string]*user.User, error) {
	if cmd.UserID == "" {
		return nil, nil, core.New(core.AuthError, "user id missing")
	}

	comps, err := h.CompRepo.FindByUser(cmd.UserID)
	if err != nil {
		return nil, nil, err
	}

	allUsers := map[string]*user.User{}
	for _, c := range comps {
		for uid := range c.Participants {
			if _, exists := allUsers[uid]; !exists {
				u, err := h.UserRepo.Get(uid)
				if err == nil {
					allUsers[uid] = u
				}
			}
		}
	}

	return comps, allUsers, nil
}
