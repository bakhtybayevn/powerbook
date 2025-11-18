package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	appUser "github.com/bakhtybayevn/powerbook/internal/application/user"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

// LoginUser godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login data"
// @Success 200 {object} dto.LoginResponse
// @Router /users/login [post]
func LoginUser(handler *appUser.LoginUserHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(core.New(core.ValidationError, err.Error()))
			return
		}

		cmd := appUser.LoginUserCommand{
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := handler.Handle(cmd)
		if err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, dto.LoginResponse{Token: token})
	}
}
