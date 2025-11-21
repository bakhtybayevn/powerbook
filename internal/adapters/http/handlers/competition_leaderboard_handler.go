package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

type LeaderboardHandler struct {
	LB ports.LeaderboardPort
}

func NewLeaderboardHandler(lb ports.LeaderboardPort) *LeaderboardHandler {
	return &LeaderboardHandler{LB: lb}
}

// GetLeaderboard godoc
// @Summary Get competition leaderboard
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Param id path string true "Competition ID"
// @Param limit query int false "Top N (default 50)"
// @Success 200 {array} map[string]interface{}
// @Router /api/v1/competitions/{id}/leaderboard [get]
func (h *LeaderboardHandler) GetLeaderboard(c *gin.Context) {
	competitionID := c.Param("id")
	if competitionID == "" {
		c.Error(core.New(core.ValidationError, "competition id is required"))
		return
	}

	limit := 50
	if q := c.Query("limit"); q != "" {
		if n, err := strconv.Atoi(q); err == nil {
			limit = n
		}
	}

	rows, err := h.LB.GetTop(c, competitionID, limit)
	if err != nil {
		c.Error(err)
		return
	}

	out := make([]gin.H, 0, len(rows))
	for i, r := range rows {
		out = append(out, gin.H{
			"rank":    i + 1,
			"user_id": r.UserID,
			"points":  r.Score,
		})
	}

	response.JSON(c, out)
}
