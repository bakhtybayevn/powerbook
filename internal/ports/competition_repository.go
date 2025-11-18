package ports

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
)

type CompetitionRepository interface {
	Save(c *competition.Competition) error
	Get(id string) (*competition.Competition, error)
	FindActive(at time.Time) ([]*competition.Competition, error)
}
