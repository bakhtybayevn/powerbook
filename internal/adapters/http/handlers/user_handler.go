package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"

    "github.com/bakhtybayevn/powerbook/internal/adapters/http/dto"
    appUser "github.com/bakhtybayevn/powerbook/internal/application/user"
)

// RegisterUser godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "User info"
// @Success 200 {object} dto.RegisterUserResponse
// @Router /users/register [post]
func RegisterUser(handler *appUser.RegisterUserHandler) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req dto.RegisterUserRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        cmd := appUser.RegisterUserCommand{
            Email:       req.Email,
            DisplayName: req.DisplayName,
            Password:    req.Password,
        }

        u, err := handler.Handle(cmd)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, dto.RegisterUserResponse{
            ID:          u.ID,
            Email:       u.Email,
            DisplayName: u.DisplayName,
        })
    }
}
