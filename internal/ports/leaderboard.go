package ports

import "context"

type LeaderboardEntry struct {
	UserID string
	Score  float64
}

type LeaderboardPort interface {
	AddScore(ctx context.Context, competitionID string, userID string, delta float64) (float64, error)
	GetTop(ctx context.Context, competitionID string, limit int) ([]LeaderboardEntry, error)
	GetRank(ctx context.Context, competitionID string, userID string) (rank int64, score float64, err error)
}

type LeaderboardHealthPort interface {
	Ping(ctx context.Context) error
}
