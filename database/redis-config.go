package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
