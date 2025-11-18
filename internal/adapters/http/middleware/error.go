package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		switch e := err.Err.(type) {

		case *core.AppError:
			switch e.Type {
			case core.ValidationError:
				response.JSONError(c, http.StatusBadRequest, e.Message)
				return
			case core.AuthError:
				response.JSONError(c, http.StatusUnauthorized, e.Message)
				return
			case core.NotFoundError:
				response.JSONError(c, http.StatusNotFound, e.Message)
				return
			default:
				response.JSONError(c, http.StatusInternalServerError, e.Message)
				return
			}

		default:
			response.JSONError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
