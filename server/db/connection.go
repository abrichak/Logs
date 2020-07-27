package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

func InitRedis() *redis.Client  {
	addr := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
