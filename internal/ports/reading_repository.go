package ports

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/reading"
)

// ReadingRepository is a port for persisting reading entries
type ReadingRepository interface {
	Save(r *reading.Reading) error
	ListByUser(userID string) ([]reading.Reading, error)
	ListByDateRange(userID string, from, to time.Time) ([]reading.Reading, error)
}
