package helpers

import (
	"github.com/gin-gonic/gin"
)

func get_content_type(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
