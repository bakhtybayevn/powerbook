package reading

import (
	"time"

	"github.com/google/uuid"
)

// Reading — value object for a single reading entry
type Reading struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Minutes   int       `json:"minutes"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
}

func NewReading(userID string, minutes int, source string, timestamp time.Time) *Reading {
	return &Reading{
		ID:        uuid.New().String(),
		UserID:    userID,
		Minutes:   minutes,
		Source:    source,
		Timestamp: timestamp,
	}
}
