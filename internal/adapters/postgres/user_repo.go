package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/user"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

// ========================================
// Save user (insert or update)
// ========================================
func (r *PostgresUserRepo) Save(u *user.User) error {
	const q = `
	INSERT INTO users (id, email, display_name, password_hash, 
	    streak_current_days, streak_last_date, total_minutes, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,NOW(),NOW())
	ON CONFLICT (id) DO UPDATE SET
	    email = EXCLUDED.email,
	    display_name = EXCLUDED.display_name,
	    password_hash = EXCLUDED.password_hash,
	    streak_current_days = EXCLUDED.streak_current_days,
	    streak_last_date = EXCLUDED.streak_last_date,
	    total_minutes = EXCLUDED.total_minutes,
	    updated_at = NOW();
	`

	_, err := r.db.Exec(q,
		u.ID,
		u.Email,
		u.DisplayName,
		u.PasswordHash,
		u.StreakCurrentDays,
		u.StreakLastDate,
		u.TotalMinutes,
	)

	if err != nil {
		log.Printf("[PostgresUserRepo.Save] SQL ERROR: %v", err)
		log.Printf("[PostgresUserRepo.Save] QUERY: %s", q)
		log.Printf("[PostgresUserRepo.Save] PARAMS: id=%s email=%s displayName=%s",
			u.ID, u.Email, u.DisplayName)

		return core.New(core.ServerError, "failed to save user")
	}
	return nil
}

// ========================================
// Get user by ID
// ========================================
func (r *PostgresUserRepo) Get(id string) (*user.User, error) {
	const q = `
	SELECT id, email, display_name, password_hash,
	       streak_current_days, streak_last_date, total_minutes
	FROM users
	WHERE id = $1;
	`

	row := r.db.QueryRow(q, id)

	var (
		u              user.User
		streakLastDate *time.Time
	)

	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.DisplayName,
		&u.PasswordHash,
		&u.StreakCurrentDays,
		&streakLastDate,
		&u.TotalMinutes,
	)

	// null → zero
	if streakLastDate != nil {
		u.StreakLastDate = streakLastDate
	}

	if err == sql.ErrNoRows {
		return nil, core.New(core.NotFoundError, "user not found")
	}

	if err != nil {
		return nil, core.New(core.ServerError, "failed to get user")
	}

	return &u, nil
}

// ========================================
// Find user by email
// ========================================
func (r *PostgresUserRepo) FindByEmail(email string) (*user.User, error) {
	const q = `
	SELECT id, email, display_name, password_hash,
	       streak_current_days, streak_last_date, total_minutes
	FROM users
	WHERE email = $1;
	`

	row := r.db.QueryRow(q, email)

	var (
		u              user.User
		streakLastDate *time.Time
	)

	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.DisplayName,
		&u.PasswordHash,
		&u.StreakCurrentDays,
		&streakLastDate,
		&u.TotalMinutes,
	)

	if streakLastDate != nil {
		u.StreakLastDate = streakLastDate
	}

	if err == sql.ErrNoRows {
		return nil, core.New(core.NotFoundError, "user not found")
	}

	if err != nil {
		return nil, core.New(core.ServerError, "failed to fetch user")
	}

	return &u, nil
}

// ========================================
// Check if email exists
// ========================================
func (r *PostgresUserRepo) Exists(email string) (bool, error) {
	const q = `SELECT 1 FROM users WHERE email = $1 LIMIT 1;`

	row := r.db.QueryRow(q, email)
	var tmp int
	err := row.Scan(&tmp)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, core.New(core.ServerError, "failed to check user existence")
	}

	return true, nil
}
