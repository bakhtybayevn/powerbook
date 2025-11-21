package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// GetRank godoc
// @Summary Get user rank in competition leaderboard
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Param id path string true "Competition ID"
// @Param userID path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /competitions/{id}/rank/{userID} [get]
func (h *LeaderboardHandler) GetRank(c *gin.Context) {
	cmpID := c.Param("id")
	userID := c.Param("userID")

	if cmpID == "" || userID == "" {
		c.Error(core.New(core.ValidationError, "missing parameters"))
		return
	}

	rank, score, err := h.LB.GetRank(c, cmpID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	if rank < 0 {
		response.JSON(c, gin.H{
			"found": false,
		})
		return
	}

	response.JSON(c, gin.H{
		"found":   true,
		"user_id": userID,
		"rank":    rank + 1,
		"points":  score,
	})
}

// GetRankMe godoc
// @Summary Get current user's rank in competition
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Param id path string true "Competition ID"
// @Success 200 {object} map[string]interface{}
// @Router /competitions/{id}/rank/me [get]
func (h *LeaderboardHandler) GetRankMe(c *gin.Context) {
	competitionID := c.Param("id")
	if competitionID == "" {
		c.Error(core.New(core.ValidationError, "competition id is required"))
		return
	}

	userID := middleware.GetUserID(c)
	if userID == "" {
		c.Error(core.New(core.AuthError, "unauthorized"))
		return
	}

	rank, score, err := h.LB.GetRank(c, competitionID, userID)
	if err != nil {
		c.Error(err)
		return
	}

	// user is not in leaderboard, not joined competition, or has 0 points
	if rank < 0 {
		response.JSON(c, gin.H{
			"found":   false,
			"user_id": userID,
		})
		return
	}

	response.JSON(c, gin.H{
		"found":   true,
		"user_id": userID,
		"rank":    rank + 1,
		"points":  score,
	})
}
