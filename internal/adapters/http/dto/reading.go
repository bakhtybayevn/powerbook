package dto

import "time"

type LogReadingRequest struct {
	Minutes   int        `json:"minutes" example:"20"`
	Source    string     `json:"source" example:"web"` // allowed: web, app, tg
	Timestamp *time.Time `json:"timestamp,omitempty" swaggertype:"string" example:"2025-11-18T12:34:56Z"`
}

type LogReadingResponse struct {
	NewStreak          int `json:"new_streak" example:"3"`
	TotalMinutesLogged int `json:"total_minutes_logged" example:"320"`
}
