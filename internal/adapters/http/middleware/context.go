package middleware

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) string {
    v, ok := c.Get("userID")
    if !ok {
        return ""
    }
    return v.(string)
}
