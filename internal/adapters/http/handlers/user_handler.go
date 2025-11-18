package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appUser "github.com/bakhtybayevn/powerbook/internal/application/user"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// RegisterUser godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "User info"
// @Success 200 {object} dto.RegisterUserResponse
// @Router /api/v1/users/register [post]
func RegisterUser(handler *appUser.RegisterUserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegisterUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(core.New(core.ValidationError, err.Error()))
			return
		}

		cmd := appUser.RegisterUserCommand{
			Email:       req.Email,
			DisplayName: req.DisplayName,
			Password:    req.Password,
		}

		u, err := handler.Handle(cmd)
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, dto.RegisterUserResponse{
			ID:          u.ID,
			Email:       u.Email,
			DisplayName: u.DisplayName,
		})
	}
}
