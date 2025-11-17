package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
	appUser "github.com/bakhtybayevn/powerbook/internal/application/user"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cmd := appUser.LoginUserCommand{
			Email:    req.Email,
			Password: req.Password,
		}

		token, err := handler.Handle(cmd)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
	}
}
