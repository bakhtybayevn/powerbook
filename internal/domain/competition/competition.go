package competition

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
)

type Competition struct {
	ID           string
	Name         string
	StartDate    time.Time
	EndDate      time.Time
	Rules        Rules
	Status       Status
	Participants map[string]*Participant // key: userID
}

func NewCompetition(name string, start, end time.Time, rules Rules) (*Competition, error) {
	if end.Before(start) {
		return nil, errors.New("competition end date cannot be before start date")
	}

	return &Competition{
		ID:           uuid.New().String(),
		Name:         name,
		StartDate:    start.UTC(),
		EndDate:      end.UTC(),
		Rules:        rules,
		Status:       StatusOpen,
		Participants: make(map[string]*Participant),
	}, nil
}

func (c *Competition) IsActive(date time.Time) bool {
	t := date.UTC()
	return c.Status == StatusOpen &&
		!t.Before(c.StartDate) &&
		!t.After(c.EndDate)
}
