package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appCompetition "github.com/bakhtybayevn/powerbook/internal/application/competition"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// CloseCompetition godoc
// @Summary Close competition and compute winners
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Param id path string true "Competition ID"
// @Success 200 {object} map[string]interface{}
// @Router /competitions/{id}/close [post]
func CloseCompetition(handler *appCompetition.CloseCompetitionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		competitionID := c.Param("id")
		if competitionID == "" {
			c.Error(core.New(core.ValidationError, "competition id is required"))
			return
		}

		winners, err := handler.Handle(appCompetition.CloseCompetitionCommand{
			CompetitionID: competitionID,
		})
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{
			"closed":  "ok",
			"winners": winners,
		})
	}
}
