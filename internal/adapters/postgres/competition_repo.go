package postgres

import (
	"database/sql"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/competition"
)

type PostgresCompetitionRepo struct {
	db *sql.DB
}

func NewPostgresCompetitionRepo(db *sql.DB) *PostgresCompetitionRepo {
	return &PostgresCompetitionRepo{db: db}
}

// --------------------------------------------------
// CREATE
// --------------------------------------------------
func (r *PostgresCompetitionRepo) Create(c *competition.Competition) error {
	const q = `
	INSERT INTO competitions (id, name, start_date, end_date, status, points_per_minute, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,NOW(),NOW());
	`

	_, err := r.db.Exec(q,
		c.ID,
		c.Name,
		c.StartDate,
		c.EndDate,
		c.Status,
		c.Rules.PointsPerMinute,
	)

	if err != nil {
		return core.New(core.ServerError, "failed to create competition")
	}
	return nil
}

// --------------------------------------------------
// GET
// --------------------------------------------------
func (r *PostgresCompetitionRepo) Get(id string) (*competition.Competition, error) {
	const compQ = `
	SELECT id, name, start_date, end_date, status, points_per_minute
	FROM competitions
	WHERE id = $1;
	`

	row := r.db.QueryRow(compQ, id)

	var c competition.Competition
	var ppm int

	err := row.Scan(
		&c.ID, &c.Name, &c.StartDate, &c.EndDate, &c.Status, &ppm,
	)

	if err == sql.ErrNoRows {
		return nil, core.New(core.NotFoundError, "competition not found")
	}
	if err != nil {
		return nil, core.New(core.ServerError, "failed to load competition")
	}

	c.Rules.PointsPerMinute = ppm

	// Load participants
	const pQ = `
	SELECT user_id, points, days_read, minutes_total, last_log_date
	FROM participants
	WHERE competition_id = $1;
	`

	rows, err := r.db.Query(pQ, id)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to load participants")
	}
	defer rows.Close()

	c.Participants = map[string]*competition.Participant{}

	for rows.Next() {
		var p competition.Participant
		var lastDate *time.Time

		err = rows.Scan(&p.UserID, &p.Points, &p.DaysRead, &p.MinutesTotal, &lastDate)
		if err != nil {
			return nil, core.New(core.ServerError, "failed to scan participant")
		}

		if lastDate != nil {
			p.LastLogDate = lastDate
		}

		c.Participants[p.UserID] = &p
	}

	return &c, nil
}

// --------------------------------------------------
// SAVE COMPETITION (update)
// --------------------------------------------------
func (r *PostgresCompetitionRepo) Save(c *competition.Competition) error {
	const q = `
	UPDATE competitions
	SET name = $2,
	    start_date = $3,
	    end_date = $4,
	    status = $5,
	    points_per_minute = $6,
	    updated_at = NOW()
	WHERE id = $1;
	`

	_, err := r.db.Exec(q,
		c.ID,
		c.Name,
		c.StartDate,
		c.EndDate,
		c.Status,
		c.Rules.PointsPerMinute,
	)

	if err != nil {
		return core.New(core.ServerError, "failed to save competition")
	}
	return nil
}

// --------------------------------------------------
// ADD / UPDATE PARTICIPANT
// --------------------------------------------------
func (r *PostgresCompetitionRepo) SaveParticipant(cID string, p *competition.Participant) error {
	const q = `
	INSERT INTO participants (competition_id, user_id, points, days_read, minutes_total, last_log_date)
	VALUES ($1,$2,$3,$4,$5,$6)
	ON CONFLICT (competition_id, user_id) DO UPDATE SET
	    points = EXCLUDED.points,
	    days_read = EXCLUDED.days_read,
	    minutes_total = EXCLUDED.minutes_total,
	    last_log_date = EXCLUDED.last_log_date;
	`

	_, err := r.db.Exec(q,
		cID,
		p.UserID,
		p.Points,
		p.DaysRead,
		p.MinutesTotal,
		p.LastLogDate,
	)
	if err != nil {
		return core.New(core.ServerError, "failed to save participant")
	}
	return nil
}

// --------------------------------------------------
// FIND ACTIVE COMPETITIONS
// --------------------------------------------------
func (r *PostgresCompetitionRepo) FindActive(ts time.Time) ([]*competition.Competition, error) {
	const q = `
	SELECT id
	FROM competitions
	WHERE start_date <= $1
	  AND end_date >= $1
	  AND status = 'open';
	`

	rows, err := r.db.Query(q, ts)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to query active competitions")
	}
	defer rows.Close()

	var list []*competition.Competition
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			continue
		}

		c, err := r.Get(id)
		if err == nil {
			list = append(list, c)
		}
	}

	return list, nil
}

// --------------------------------------------------
// GET ALL COMPETITIONS
// --------------------------------------------------
func (r *PostgresCompetitionRepo) GetAll() ([]*competition.Competition, error) {
	const q = `
        SELECT id
        FROM competitions
        ORDER BY start_date DESC;
    `

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to load competitions")
	}
	defer rows.Close()

	var list []*competition.Competition
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			continue
		}

		c, err := r.Get(id)
		if err == nil {
			list = append(list, c)
		}
	}

	return list, nil
}

// --------------------------------------------------
// GET COMPETITIONS WHERE USER PARTICIPATES
// --------------------------------------------------
func (r *PostgresCompetitionRepo) FindByUser(userID string) ([]*competition.Competition, error) {
	const q = `
        SELECT competition_id
        FROM participants
        WHERE user_id = $1;
    `

	rows, err := r.db.Query(q, userID)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to load user's competitions")
	}
	defer rows.Close()

	var list []*competition.Competition
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			continue
		}

		c, err := r.Get(id)
		if err == nil {
			list = append(list, c)
		}
	}

	return list, nil
}
