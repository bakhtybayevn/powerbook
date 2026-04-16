package competition

import (
	"time"

	"github.com/google/uuid"
)

// GiftExchange represents a gift pairing between top 50% and bottom 50% after competition ends.
type GiftExchange struct {
	ID                string
	CompetitionID     string
	GiverID           string
	ReceiverID        string
	GiftDescription   string
	GiverConfirmed    bool
	ReceiverConfirmed bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewGiftExchange(competitionID, giverID, receiverID string) *GiftExchange {
	return &GiftExchange{
		ID:            uuid.New().String(),
		CompetitionID: competitionID,
		GiverID:       giverID,
		ReceiverID:    receiverID,
	}
}
