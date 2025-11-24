package ports

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
)

type CompetitionRepository interface {
	Create(c *competition.Competition) error
	Save(c *competition.Competition) error
	SaveParticipant(competitionID string, p *competition.Participant) error
	Get(id string) (*competition.Competition, error)
	FindActive(at time.Time) ([]*competition.Competition, error)
	GetAll() ([]*competition.Competition, error)
	FindByUser(userID string) ([]*competition.Competition, error)
}
