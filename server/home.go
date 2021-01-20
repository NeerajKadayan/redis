package server

import (
	"github.com/gin-gonic/gin"
)

// Test --
func Home(c *gin.Context) {
	c.String(200, "REDIS MONITOR")
}
