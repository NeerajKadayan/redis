package store

import (
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

// connect -- initiates connection to redis
func ConnectToRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     5,
		MinIdleConns: 2,
	})

	_, err := Rdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
