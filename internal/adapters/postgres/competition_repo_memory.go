package postgres

import (
	"errors"
	"sync"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
)

type InMemoryCompetitionRepo struct {
	mu    sync.RWMutex
	store map[string]*competition.Competition
}

func NewInMemoryCompetitionRepo() *InMemoryCompetitionRepo {
	return &InMemoryCompetitionRepo{
		store: make(map[string]*competition.Competition),
	}
}

func (r *InMemoryCompetitionRepo) Save(c *competition.Competition) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[c.ID] = c
	return nil
}

func (r *InMemoryCompetitionRepo) Get(id string) (*competition.Competition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.store[id]
	if !ok {
		return nil, errors.New("competition not found")
	}
	return c, nil
}

func (r *InMemoryCompetitionRepo) FindActive(at time.Time) ([]*competition.Competition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := []*competition.Competition{}

	for _, cmp := range r.store {
		if cmp.IsActive(at) {
			list = append(list, cmp)
		}
	}

	return list, nil
}
