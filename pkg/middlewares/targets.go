package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
)

func TargetsMiddleware(targets *pkg.Targets) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(pkg.STargets, targets)
		c.Next()
	}
}
