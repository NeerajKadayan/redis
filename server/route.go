package server

import (
	"github.com/gin-gonic/gin"
)

// RouterMain --
func RouterMain(route *gin.Engine) {

	route.GET("/", Home)

	// web socket
	route.GET("/ws", WebSocketAPI)
	route.GET("/redis", RedisInfo)

}
