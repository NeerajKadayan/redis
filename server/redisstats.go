package server

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/neeraj-singh/redis/store"
	"github.com/neeraj-singh/redis/types"
)

// RedisInfo --
func RedisInfo(c *gin.Context) {
	r := store.Rdb
	var inputKeys []string
	inputKeys = append(inputKeys, "CPU", "CLUSTER", "CLIENTS", "MEMORY", "SERVER")
	if len(inputKeys) == 0 {
		c.JSON(400, gin.H{"message": "no input keys found!"})
	}

	resp := processKeys(r, inputKeys)
	c.String(200, "utf-8", resp)
}

// processKeys --
func processKeys(r *redis.Client, input []string) (val types.Stats) {
	var wg sync.WaitGroup
	wg.Add(len(input))
	listener := make(chan types.Stats, len(input))
	values := make(types.Stats, 0)

	for i := range input {
		go GetStats(input[i], r, listener, values, &wg)
		val = <-listener
	}

	wg.Wait()
	close(listener)
	return
}

// GetStats --
func GetStats(input string, r *redis.Client, l chan types.Stats,
	val types.Stats, wg *sync.WaitGroup) {

	defer wg.Done()

	switch {

	case input == "CPU":
		v := r.Info("CPU")
		val["CPU"] = v
		l <- val

	case input == "CLIENTS":
		v := r.Info("CLIENTS")
		val["CLIENTS"] = v
		l <- val

	case input == "CLUSTER":
		v := r.Info("CLUSTER")
		val["CLUSTER"] = v
		l <- val

	case input == "SERVER":
		v := r.Info("SERVER")
		val["SERVER"] = v
		l <- val

	case input == "MEMORY":
		v := r.Info("MEMORY")
		val["MEMORY"] = v
		l <- val

	}
}
