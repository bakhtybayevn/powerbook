package ports

import "github.com/bakhtybayevn/powerbook/internal/domain/user"

type UserRepository interface {
	Get(id string) (*user.User, error)
	Save(u *user.User) error
	FindByEmail(email string) (*user.User, error)
}
