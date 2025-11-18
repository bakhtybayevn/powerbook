package ports

import "github.com/bakhtybayevn/powerbook/internal/domain/reading"

// ReadingRepository is a port for persisting reading entries
type ReadingRepository interface {
	Save(r *reading.Reading) error
	// optionally: ListByUser(userID string) ([]*reading.Reading, error)
}
