package server

import (
	"github.com/company-tests/tazapay/test/websocket"
	"github.com/gin-gonic/gin"
)

// WebSocketAPI --
func WebSocketAPI(c *gin.Context) {
	websocket.Melody()
	c.JSON(200, gin.H{"message": "this is a message from socket API"})
}
