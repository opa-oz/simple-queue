package checks

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

var RedisNotWorking = errors.New("Redis is not working")

func CheckRedis(c context.Context, rdb *redis.Client) error {
	echo := rdb.Ping(c)

	if echo.Err() != nil {
		return RedisNotWorking
	}

	return nil
}
