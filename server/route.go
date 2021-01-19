package server

import (
	"github.com/company-tests/tazapay/test/websocket"
	"github.com/gin-gonic/gin"
)

// RouterMain --
func RouterMain(route *gin.Engine) {

	ws := route.Group("/ws")
	ws.GET("/test", func(c *gin.Context) {
		// web socket
		websocket.WShandler(c.Writer, c.Request)

	})

}
