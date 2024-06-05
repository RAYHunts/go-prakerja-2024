package middleware

import (
	"net/http"
	"strings"
	"uk/helper"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		token = splitToken[1]

		if !helper.ValidateToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"token":   token,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

