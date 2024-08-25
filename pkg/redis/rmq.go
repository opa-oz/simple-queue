package redis

import (
	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

func GetRMQConnection(rdb *redis.Client, errChan chan<- error) (*rmq.Connection, error) {
	conn, err := rmq.OpenConnectionWithRedisClient("Simple queue", rdb, errChan)

	return &conn, err
}
