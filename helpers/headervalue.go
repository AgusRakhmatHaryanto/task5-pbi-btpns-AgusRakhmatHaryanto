package helpers

import (
	"github.com/gin-gonic/gin"
)

func HeaderValue(c *gin.Context) string {

	value := c.Request.Header.Get("Content-Type")
	return value
}