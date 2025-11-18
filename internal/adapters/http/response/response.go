package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ======================
// SUCCESS RESPONSE
// ======================

type Success struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Success{
		Success: true,
		Data:    data,
	})
}

// ======================
// ERROR RESPONSE
// ======================

type Error struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func JSONError(c *gin.Context, status int, msg string) {
	c.AbortWithStatusJSON(status, Error{
		Success: false,
		Error:   msg,
	})
}
