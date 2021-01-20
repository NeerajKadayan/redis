package server

import (
	"github.com/gin-gonic/gin"
)

// Test --
func Test(c *gin.Context) {
	c.String(200, "We got Gin")
}
