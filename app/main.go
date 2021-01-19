package main

import (
	"log"

	"github.com/company-tests/tazapay/test/server"
	"github.com/company-tests/tazapay/test/store"
	"github.com/gin-gonic/gin"
)

func main() {
	err := store.ConnectToRedis()
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.String(200, "We got Gin")
	})

	server.RouterMain(app)

	err = app.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
