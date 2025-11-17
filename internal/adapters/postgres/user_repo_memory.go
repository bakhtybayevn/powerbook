package postgres

import (
	"errors"

	"github.com/bakhtybayevn/powerbook/internal/domain/user"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type InMemoryUserRepo struct {
	users map[string]*user.User
}

func NewInMemoryUserRepo() ports.UserRepository {
	return &InMemoryUserRepo{
		users: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepo) Get(id string) (*user.User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepo) Save(u *user.User) error {
	r.users[u.ID] = u
	return nil
}

func (r *InMemoryUserRepo) FindByEmail(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
