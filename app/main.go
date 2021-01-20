package main

import (
	"log"

	"github.com/company-tests/tazapay/test/server"
	"github.com/company-tests/tazapay/test/store"
	"github.com/gin-gonic/gin"
)

func main() {
	// establish connection to redis
	err := store.ConnectToRedis()
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	server.RouterMain(app)

	err = app.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

}
