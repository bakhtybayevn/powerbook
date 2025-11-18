package user

import (
	"net/mail"
	"strings"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/user"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type RegisterUserCommand struct {
	Email       string
	DisplayName string
	Password    string
}

type RegisterUserHandler struct {
	Repo ports.UserRepository
}

func NewRegisterUserHandler(repo ports.UserRepository) *RegisterUserHandler {
	return &RegisterUserHandler{Repo: repo}
}

func (h *RegisterUserHandler) Handle(cmd RegisterUserCommand) (*user.User, error) {
	// === VALIDATION ===

	// email
	cmd.Email = strings.TrimSpace(cmd.Email)
	if cmd.Email == "" {
		return nil, core.New(core.ValidationError, "email is required")
	}
	if _, err := mail.ParseAddress(cmd.Email); err != nil {
		return nil, core.New(core.ValidationError, "invalid email format")
	}

	// display name
	if len(cmd.DisplayName) < 2 {
		return nil, core.New(core.ValidationError, "display name must be at least 2 characters")
	}
	if len(cmd.DisplayName) > 64 {
		return nil, core.New(core.ValidationError, "display name too long")
	}

	// password
	if len(cmd.Password) < 6 {
		return nil, core.New(core.ValidationError, "password must be at least 6 characters")
	}
	if len(cmd.Password) > 64 {
		return nil, core.New(core.ValidationError, "password too long")
	}

	// unique email check
	exist, _ := h.Repo.FindByEmail(cmd.Email)
	if exist != nil {
		return nil, core.New(core.ValidationError, "user with this email already exists")
	}

	u := user.NewUser(cmd.Email, cmd.DisplayName, cmd.Password)
	if err := h.Repo.Save(u); err != nil {
		return nil, core.Wrap(err, core.ServerError)
	}

	return u, nil
}
