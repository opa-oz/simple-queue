package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
	redis2 "github.com/redis/go-redis/v9"
)

func RedisMiddleware(rdb *redis2.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(pkg.SRedis, rdb)
		c.Next()
	}
}
