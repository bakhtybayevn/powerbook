package middleware

import (
	"net/http"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/ports"
	"github.com/gin-gonic/gin"
)

func RedisHealth(redis ports.LeaderboardHealthPort) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := redis.Ping(c); err != nil {
			response.JSONError(c, http.StatusServiceUnavailable, "leaderboard service unavailable")
			c.Abort()
			return
		}
		c.Next()
	}
}
