package server

import (
	"github.com/gin-gonic/gin"
	"github.com/neeraj-singh/redis/websocket"
)

// WebSocketAPI --
func WebSocketAPI(c *gin.Context) {
	websocket.Melody()
	c.JSON(200, gin.H{"message": "this is a message from socket API"})
}
