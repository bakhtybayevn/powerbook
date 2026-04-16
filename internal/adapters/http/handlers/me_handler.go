package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

// GetMe godoc
// @Summary Get authenticated user info with stats
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/me [get]
func GetMe(repo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		u, err := repo.Get(userID)
		if err != nil {
			c.Error(core.New(core.NotFoundError, "user not found"))
			return
		}

		response.JSON(c, gin.H{
			"id":              u.ID,
			"email":           u.Email,
			"display_name":    u.DisplayName,
			"streak_current":  u.StreakCurrentDays,
			"total_minutes":   u.TotalMinutes,
			"xp":              u.XP,
			"level":           u.Level(),
			"level_name":      u.LevelName(),
			"telegram_handle": u.TelegramHandle,
			"is_admin":        u.IsAdmin,
		})
	}
}
