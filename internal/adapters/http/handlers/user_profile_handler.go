package handlers

import (
	"github.com/gin-gonic/gin"

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
			"id":             u.ID,
			"display_name":   u.DisplayName,
			"streak_current": u.StreakCurrentDays,
			"total_minutes":  u.TotalMinutes,
		})
	}
}
