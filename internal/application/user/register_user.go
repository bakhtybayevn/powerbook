package user

import (
    "errors"
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
    if cmd.Email == "" || cmd.Password == "" {
        return nil, errors.New("email and password required")
    }

    newUser := user.NewUser(cmd.Email, cmd.DisplayName, cmd.Password)
    if err := h.Repo.Save(newUser); err != nil {
        return nil, err
    }

    return newUser, nil
}
