package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

// GetUserProfile godoc
// @Summary Get public user profile by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUserProfile(repo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		if userID == "" {
			c.Error(core.New(core.ValidationError, "user id is required"))
			return
		}

		u, err := repo.Get(userID)
		if err != nil {
			c.Error(core.New(core.NotFoundError, "user not found"))
			return
		}

		response.JSON(c, gin.H{
			"id":              u.ID,
			"display_name":    u.DisplayName,
			"streak_current":  u.StreakCurrentDays,
			"total_minutes":   u.TotalMinutes,
			"xp":              u.XP,
			"level":           u.Level(),
			"level_name":      u.LevelName(),
			"telegram_handle": u.TelegramHandle,
		})
	}
}

// UpdateProfile godoc
// @Summary Update current user's profile (telegram handle)
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/me/profile [put]
func UpdateProfile(repo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.Error(core.New(core.AuthError, "unauthorized"))
			return
		}

		var req struct {
			TelegramHandle string `json:"telegram_handle"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(core.New(core.ValidationError, "invalid request body"))
			return
		}

		u, err := repo.Get(userID)
		if err != nil {
			c.Error(core.New(core.NotFoundError, "user not found"))
			return
		}

		u.TelegramHandle = req.TelegramHandle
		if err := repo.Save(u); err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{
			"id":              u.ID,
			"display_name":    u.DisplayName,
			"telegram_handle": u.TelegramHandle,
		})
	}
}
