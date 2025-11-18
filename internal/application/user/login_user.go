package user

import (
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type LoginUserCommand struct {
	Email    string
	Password string
}

type LoginUserHandler struct {
	Repo  ports.UserRepository
	Token ports.TokenService
}

func NewLoginUserHandler(repo ports.UserRepository, token ports.TokenService) *LoginUserHandler {
	return &LoginUserHandler{Repo: repo, Token: token}
}

func (h *LoginUserHandler) Handle(cmd LoginUserCommand) (string, error) {
	// Поиск пользователя по email
	u, err := h.Repo.FindByEmail(cmd.Email)
	if err != nil {
		return "", core.New(core.ValidationError, "invalid email")
	}

	if !u.CheckPassword(cmd.Password) {
		return "", core.New(core.ValidationError, "invalid password")
	}

	return h.Token.GenerateToken(u.ID)
}
