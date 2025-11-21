package redis

import (
	"context"
	"fmt"

	"github.com/bakhtybayevn/powerbook/internal/ports"
	"github.com/redis/go-redis/v9"
)

type RedisLeaderboard struct {
	client *redis.Client
}

func NewRedisLeaderboard(addr string) *RedisLeaderboard {
	return &RedisLeaderboard{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   0,
		}),
	}
}

func (r *RedisLeaderboard) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisLeaderboard) key(competitionID string) string {
	return fmt.Sprintf("leaderboard:cmp:%s", competitionID)
}

func (r *RedisLeaderboard) AddScore(ctx context.Context, competitionID string, userID string, delta float64) (float64, error) {
	return r.client.ZIncrBy(ctx, r.key(competitionID), delta, userID).Result()
}

func (r *RedisLeaderboard) GetTop(ctx context.Context, competitionID string, limit int) ([]ports.LeaderboardEntry, error) {
	results, err := r.client.ZRevRangeWithScores(ctx, r.key(competitionID), 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	out := make([]ports.LeaderboardEntry, 0, len(results))
	for _, z := range results {
		out = append(out, ports.LeaderboardEntry{
			UserID: z.Member.(string),
			Score:  z.Score,
		})
	}

	return out, nil
}

func (r *RedisLeaderboard) GetRank(ctx context.Context, competitionID string, userID string) (int64, float64, error) {
	rank, err := r.client.ZRevRank(ctx, r.key(competitionID), userID).Result()
	if err == redis.Nil {
		return -1, 0, nil
	}
	if err != nil {
		return -1, 0, err
	}

	score, err := r.client.ZScore(ctx, r.key(competitionID), userID).Result()
	if err != nil {
		return rank, 0, err
	}

	return rank, score, nil
}
