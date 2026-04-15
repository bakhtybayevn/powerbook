package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

// ReadingHistory godoc
// @Summary Get user's reading history
// @Tags reading
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /reading/history [get]
func ReadingHistory(repo ports.ReadingRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)

		logs, err := repo.ListByUser(userID)
		if err != nil {
			c.Error(core.New(core.ServerError, "failed to load reading history"))
			return
		}

		type entry struct {
			ID        string `json:"id"`
			Minutes   int    `json:"minutes"`
			Source    string `json:"source"`
			Timestamp string `json:"timestamp"`
		}

		entries := make([]entry, 0, len(logs))
		for _, l := range logs {
			entries = append(entries, entry{
				ID:        l.ID,
				Minutes:   l.Minutes,
				Source:    l.Source,
				Timestamp: l.Timestamp.Format("2006-01-02T15:04:05Z"),
			})
		}

		response.JSON(c, gin.H{
			"readings": entries,
			"total":    len(entries),
		})
	}
}
