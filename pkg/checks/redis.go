package checks

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

const EchoCheck string = "I'm alive"

var RedisNotWorking = errors.New("Redis is not working")

func CheckRedis(c context.Context, rdb *redis.Client) error {
	echo := rdb.Echo(c, EchoCheck)

	if echo.Err() != nil {
		return RedisNotWorking
	}

	return nil
}
