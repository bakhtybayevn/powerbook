package postgres

import (
	"errors"
	"sync"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/reading"
)

type InMemoryReadingRepo struct {
	mu      sync.RWMutex
	storage map[string]*reading.Reading
	lastSeq int64
}

func NewInMemoryReadingRepo() *InMemoryReadingRepo {
	return &InMemoryReadingRepo{
		storage: make(map[string]*reading.Reading),
	}
}

func (r *InMemoryReadingRepo) nextID() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lastSeq++
	return "reading-" + time.Now().Format("20060102T150405") + "-" + string(r.lastSeq)
}

func (r *InMemoryReadingRepo) Save(rd *reading.Reading) error {
	if rd == nil {
		return errors.New("nil reading")
	}
	// Generate ID if empty
	if rd.ID == "" {
		rd.ID = "rid-" + time.Now().UTC().Format("20060102T150405.000000")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.storage[rd.ID] = rd
	return nil
}
