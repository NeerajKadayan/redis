package store

import (
	"fmt"

	"github.com/go-redis/redis"
)

// connect -- initiates connection to redis
func ConnectToRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     5,
		MinIdleConns: 2,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	r := client.Get("test")
	fmt.Println(r)

	r.Val()

	return nil
}
