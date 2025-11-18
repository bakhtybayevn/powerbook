package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appCompetition "github.com/bakhtybayevn/powerbook/internal/application/competition"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// CreateCompetition godoc
// @Summary Create a new competition
// @Tags competition
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateCompetitionRequest true "Competition info"
// @Success 200 {object} dto.CompetitionResponse
// @Router /api/v1/competitions/create [post]
func CreateCompetition(handler *appCompetition.CreateCompetitionHandler) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req dto.CreateCompetitionRequest

        if err := c.ShouldBindJSON(&req); err != nil {
            c.Error(core.New(core.ValidationError, "invalid competition request"))
            return
        }

        cmp, err := handler.Handle(appCompetition.CreateCompetitionCommand{
            Name:            req.Name,
            StartDate:       req.StartDate,
            EndDate:         req.EndDate,
            PointsPerMinute: req.PointsPerMinute,
        })
        if err != nil {
            c.Error(err)
            return
        }

        response.JSON(c, dto.CompetitionResponse{
            ID:              cmp.ID,
            Name:            cmp.Name,
            StartDate:       cmp.StartDate,
            EndDate:         cmp.EndDate,
            Status:          string(cmp.Status),
            PointsPerMinute: cmp.Rules.PointsPerMinute,
        })
    }
}
