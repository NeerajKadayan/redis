package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/neeraj-singh/redis/server"
	"github.com/neeraj-singh/redis/store"
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
