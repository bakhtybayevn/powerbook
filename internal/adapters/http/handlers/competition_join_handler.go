package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appCompetition "github.com/bakhtybayevn/powerbook/internal/application/competition"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// JoinCompetition godoc
// @Summary Join a competition
// @Tags competition
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Competition ID"
// @Success 200 {object} map[string]string
// @Router /competitions/{id}/join [post]
func JoinCompetition(handler *appCompetition.JoinCompetitionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.Error(core.New(core.AuthError, "unauthorized"))
			return
		}

		competitionID := c.Param("id")
		if competitionID == "" {
			c.Error(core.New(core.ValidationError, "competition id is required"))
			return
		}

		err := handler.Handle(appCompetition.JoinCompetitionCommand{
			UserID:        userID,
			CompetitionID: competitionID,
		})

		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{"joined": "ok"})
	}
}
