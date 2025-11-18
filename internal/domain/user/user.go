package user

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User aggregate
type User struct {
	ID           string
	Email        string
	DisplayName  string
	PasswordHash string

	// streak tracking
	StreakCurrentDays int
	StreakLastDate    *time.Time

	// analytics
	TotalMinutes int
}

func NewUser(email, displayName, password string) *User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		ID:                uuid.New().String(),
		Email:             email,
		DisplayName:       displayName,
		PasswordHash:      string(hash),
		StreakCurrentDays: 0,
		StreakLastDate:    nil,
		TotalMinutes:      0,
	}
}

// LogReading performs domain logic for a user's reading entry.
// minutes must be > 0. timestamp is the time of the reading event (UTC).
// Returns newStreak, totalMinutes, and an error (domain error via core.AppError should be returned by application layer if needed).
func (u *User) LogReading(minutes int, timestamp time.Time) (newStreak int, totalMinutes int) {
	// Normalize timestamp to date (UTC) for streak counting
	t := timestamp.UTC()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	// Determine previous day date
	var yesterday time.Time
	if u.StreakLastDate != nil {
		ld := u.StreakLastDate.UTC()
		yesterday = time.Date(ld.Year(), ld.Month(), ld.Day(), 0, 0, 0, 0, time.UTC)
	}

	// Update streak logic
	if u.StreakLastDate == nil {
		u.StreakCurrentDays = 1
	} else if today.Equal(yesterday) {
		// multiple logs in same day -> do not increment streak (streak counts days with at least one log)
		// but keep streak unchanged
		// Note: we still update last date to today if necessary (it already equals)
	} else {
		// if yesterday was exactly previous day -> increment, else reset to 1
		prevDay := yesterday.AddDate(0, 0, 1)
		if prevDay.Equal(today) {
			u.StreakCurrentDays += 1
		} else {
			u.StreakCurrentDays = 1
		}
	}

	// Update last log date to today
	u.StreakLastDate = &today

	// Update total minutes
	u.TotalMinutes += minutes

	return u.StreakCurrentDays, u.TotalMinutes
}

func (u *User) CheckPassword(password string) bool {
	log.Println("Checking password for user:", u.Email)
	log.Println("Stored hash:", u.PasswordHash)
	log.Println("Provided password:", password)
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) == nil
}
