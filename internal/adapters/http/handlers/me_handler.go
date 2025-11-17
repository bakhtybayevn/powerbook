package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

// GetMe godoc
// @Summary Get authenticated user info
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /users/me [get]
func GetMe(repo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		u, err := repo.Get(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":           u.ID,
			"email":        u.Email,
			"display_name": u.DisplayName,
		})
	}
}
