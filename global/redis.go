package global

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return
	}

	fmt.Println("ping:", pong)

	return
}
