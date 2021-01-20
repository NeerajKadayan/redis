package server

import (
	"github.com/gin-gonic/gin"
)

// RouterMain --
func RouterMain(route *gin.Engine) {

	route.GET("/", Test)

	// web socket
	route.GET("/ws", WebSocketAPI)
	route.GET("/redis", RedisInfo)

}
