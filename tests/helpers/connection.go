package helpers

import (
	mockRedis "github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
)

func InitRedis() (*mockRedis.Miniredis, *redis.Client)  {
	mr, err := mockRedis.Run()
	if err != nil {
		panic(err)
	}

	return mr, redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}
