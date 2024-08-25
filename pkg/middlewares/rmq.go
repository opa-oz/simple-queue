package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
)

func RMQMiddleware(queues *pkg.RMQueues) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("RMQ", queues)
		c.Next()
	}
}
