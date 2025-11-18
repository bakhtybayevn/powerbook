package reading

import "time"

// Reading — value object for a single reading entry
type Reading struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Minutes   int       `json:"minutes"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
}
