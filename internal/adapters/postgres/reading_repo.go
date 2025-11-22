package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/reading"
)

type PostgresReadingRepo struct {
	db *sql.DB
}

func NewPostgresReadingRepo(db *sql.DB) *PostgresReadingRepo {
	return &PostgresReadingRepo{db: db}
}

func (r *PostgresReadingRepo) Save(rd *reading.Reading) error {
	const q = `
	INSERT INTO reading_logs (id, user_id, minutes, source, timestamp, created_at)
	VALUES ($1, $2, $3, $4, $5, NOW());
	`

	_, err := r.db.Exec(q,
		rd.ID,
		rd.UserID,
		rd.Minutes,
		rd.Source,
		rd.Timestamp,
	)

	if err != nil {
		log.Printf("[PostgresReadingRepo.Save] SQL ERROR: %v", err)
		log.Printf("[PostgresReadingRepo.Save] QUERY: %s", q)
		log.Printf("[PostgresReadingRepo.Save] PARAMS: id=%s userID=%s minutes=%d source=%s timestamp=%v",
			rd.ID, rd.UserID, rd.Minutes, rd.Source, rd.Timestamp)
		return core.New(core.ServerError, "failed to save reading log")
	}
	return nil
}

func (r *PostgresReadingRepo) ListByUser(userID string) ([]reading.Reading, error) {
	const q = `
	SELECT id, user_id, minutes, source, timestamp
	FROM reading_logs
	WHERE user_id = $1
	ORDER BY timestamp DESC;
	`

	rows, err := r.db.Query(q, userID)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to query reading logs")
	}
	defer rows.Close()

	var list []reading.Reading

	for rows.Next() {
		var rd reading.Reading
		if err := rows.Scan(&rd.ID, &rd.UserID, &rd.Minutes, &rd.Source, &rd.Timestamp); err != nil {
			return nil, core.New(core.ServerError, "failed to scan reading log")
		}
		list = append(list, rd)
	}

	return list, nil
}

func (r *PostgresReadingRepo) ListByDateRange(userID string, from, to time.Time) ([]reading.Reading, error) {
	const q = `
	SELECT id, user_id, minutes, source, timestamp
	FROM reading_logs
	WHERE user_id = $1
	  AND timestamp BETWEEN $2 AND $3
	ORDER BY timestamp ASC;
	`

	rows, err := r.db.Query(q, userID, from, to)
	if err != nil {
		return nil, core.New(core.ServerError, "failed to query logs by date range")
	}
	defer rows.Close()

	var list []reading.Reading

	for rows.Next() {
		var rd reading.Reading
		if err := rows.Scan(&rd.ID, &rd.UserID, &rd.Minutes, &rd.Source, &rd.Timestamp); err != nil {
			return nil, core.New(core.ServerError, "failed to scan log")
		}
		list = append(list, rd)
	}

	return list, nil
}
