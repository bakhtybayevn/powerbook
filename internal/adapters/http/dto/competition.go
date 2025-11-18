package dto

import "time"

// Request to create competition
type CreateCompetitionRequest struct {
	Name            string    `json:"name" example:"January Reading Contest"`
	StartDate       time.Time `json:"start_date" example:"2025-01-01T00:00:00Z"`
	EndDate         time.Time `json:"end_date" example:"2025-01-31T23:59:59Z"`
	PointsPerMinute int       `json:"points_per_minute" example:"1"`
}

// Response for created competition
type CompetitionResponse struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          string    `json:"status"`
	PointsPerMinute int       `json:"points_per_minute"`
}

// Join request
type JoinCompetitionRequest struct {
	CompetitionID string `json:"competition_id" example:"cmp-123"`
}

// Close competition response (summary)
type CloseCompetitionResponse struct {
	Winners []string `json:"winners"`
}
