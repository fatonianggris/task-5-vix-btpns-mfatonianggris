package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.verify_tokens(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
