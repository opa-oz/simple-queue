package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/redis/go-redis/v9"
)

func GetRedis(c *gin.Context) (*redis.Client, error) {
	r := c.Value("Redis")

	if r == nil {
		err := fmt.Errorf("could not retrieve Redis")
		return nil, err
	}

	rdb, ok := r.(*redis.Client)
	if !ok {
		err := fmt.Errorf("variable Redis has wrong type")
		return nil, err
	}

	return rdb, nil
}

func GetRMQ(c *gin.Context) (*pkg.RMQueues, error) {
	r := c.Value("RMQ")

	if r == nil {
		err := fmt.Errorf("could not retrieve Redis queue")
		return nil, err
	}

	connection, ok := r.(*pkg.RMQueues)
	if !ok {
		err := fmt.Errorf("variable Redis Queue has wrong type")
		return nil, err
	}

	return connection, nil
}
