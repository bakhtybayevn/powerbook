package handlers

import (
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	response "github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/application/competition"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/gin-gonic/gin"
)

// @Summary List all competitions
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /competitions [get]
func ListAllCompetitions(h *competition.ListAllCompetitionsHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		comps, users, err := h.Handle()
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{
			"competitions": dto.CompetitionsToDTO(comps, users),
		})
	}
}

// @Summary List competitions where current user participates
// @Tags competition
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /competitions/my [get]
func ListMyCompetitions(h *competition.ListMyCompetitionsHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.Error(core.New(core.AuthError, "unauthorized"))
			return
		}

		comps, users, err := h.Handle(competition.ListMyCommand{UserID: userID})
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{
			"competitions": dto.CompetitionsToDTO(comps, users),
		})
	}
}
