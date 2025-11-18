package handlers

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appReading "github.com/bakhtybayevn/powerbook/internal/application/reading"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// LogReading godoc
// @Summary Log reading minutes for authenticated user
// @Tags reading
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.LogReadingRequest true "Reading data"
// @Success 200 {object} dto.LogReadingResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/reading/log [post]
func LogReading(handler *appReading.LogReadingHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.Error(core.New(core.AuthError, "unauthorized"))
			return
		}

		var req dto.LogReadingRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(core.New(core.ValidationError, "invalid request payload"))
			return
		}

		// use provided timestamp if present, else server time
		ts := time.Now().UTC()
		if req.Timestamp != nil {
			ts = req.Timestamp.UTC()
		}

		newStreak, totalMinutes, err := handler.Handle(appReading.LogReadingCommand{
			UserID:    userID,
			Minutes:   req.Minutes,
			Source:    req.Source,
			Timestamp: ts,
		})
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, dto.LogReadingResponse{
			NewStreak:          newStreak,
			TotalMinutesLogged: totalMinutes,
		})
	}
}
