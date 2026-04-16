package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/domain/user"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

func GetCompetition(compRepo ports.CompetitionRepository, userRepo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.Error(core.New(core.ValidationError, "competition id is required"))
			return
		}

		comp, err := compRepo.Get(id)
		if err != nil {
			c.Error(err)
			return
		}

		allUsers := map[string]*user.User{}
		for uid := range comp.Participants {
			if u, err := userRepo.Get(uid); err == nil {
				allUsers[uid] = u
			}
		}

		response.JSON(c, dto.CompetitionToDTO(comp, allUsers))
	}
}
