package competition

import "time"

type Participant struct {
	UserID       string
	Points       int
	DaysRead     int
	LastLogDate  *time.Time
	MinutesTotal int
}

func NewParticipant(userID string) *Participant {
	return &Participant{
		UserID:       userID,
		Points:       0,
		DaysRead:     0,
		MinutesTotal: 0,
	}
}

func (p *Participant) AddReading(minutes int, timestamp time.Time, rules Rules) {
	pointsEarned := rules.PointsPerMinute * minutes
	p.Points += pointsEarned

	p.MinutesTotal += minutes

	// streak inside competition
	day := timestamp.UTC().Truncate(24 * time.Hour)

	if p.LastLogDate == nil {
		p.DaysRead = 1
	} else {
		prevDay := p.LastLogDate.UTC().Truncate(24 * time.Hour)
		if day.After(prevDay) {
			p.DaysRead++
		}
	}
	p.LastLogDate = &day
}
