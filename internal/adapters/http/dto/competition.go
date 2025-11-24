package dto

import (
	"time"

	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
	"github.com/bakhtybayevn/powerbook/internal/domain/user"
)

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

type PublicUserDTO struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type ParticipantDTO struct {
	User         PublicUserDTO `json:"user"`
	Points       int           `json:"points"`
	DaysRead     int           `json:"days_read"`
	MinutesTotal int           `json:"minutes_total"`
	LastLogDate  *string       `json:"last_log_date,omitempty"`
}

type CompetitionDTO struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	StartDate    string           `json:"start_date"`
	EndDate      string           `json:"end_date"`
	Status       string           `json:"status"`
	Points       int              `json:"points_per_minute"`
	Participants []ParticipantDTO `json:"participants,omitempty"`
}

func UserToPublicDTO(u *user.User) PublicUserDTO {
	return PublicUserDTO{
		ID:          u.ID,
		DisplayName: u.DisplayName,
	}
}

func ParticipantToDTO(p *competition.Participant, u *user.User) ParticipantDTO {
	var last *string
	if p.LastLogDate != nil {
		s := p.LastLogDate.Format(time.RFC3339)
		last = &s
	}

	return ParticipantDTO{
		User:         UserToPublicDTO(u),
		Points:       p.Points,
		DaysRead:     p.DaysRead,
		MinutesTotal: p.MinutesTotal,
		LastLogDate:  last,
	}
}

func CompetitionToDTO(c *competition.Competition, users map[string]*user.User) CompetitionDTO {
	participants := make([]ParticipantDTO, 0, len(c.Participants))

	for uid, p := range c.Participants {
		u := users[uid]
		participants = append(participants, ParticipantToDTO(p, u))
	}

	return CompetitionDTO{
		ID:           c.ID,
		Name:         c.Name,
		StartDate:    c.StartDate.Format(time.RFC3339),
		EndDate:      c.EndDate.Format(time.RFC3339),
		Status:       string(c.Status),
		Points:       c.Rules.PointsPerMinute,
		Participants: participants,
	}
}

func CompetitionsToDTO(list []*competition.Competition, users map[string]*user.User) []CompetitionDTO {
	out := make([]CompetitionDTO, 0, len(list))
	for _, c := range list {
		out = append(out, CompetitionToDTO(c, users))
	}
	return out
}
