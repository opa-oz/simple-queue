package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg/api"
	"github.com/opa-oz/simple-queue/pkg/config"
	"github.com/opa-oz/simple-queue/pkg/middlewares"
	"github.com/opa-oz/simple-queue/pkg/redis"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	rdb := redis.GetClient(cfg)

	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.ResponseLogger())
	r.Use(middlewares.RedisMiddleware(rdb))
	r.Use(middlewares.CfgMiddleware(cfg))

	r.GET("/healz", api.Healz)
	r.GET("/ready", api.Ready)

	// rg := r.Group("/api")

	port := fmt.Sprintf(":%d", cfg.Port)
	err = r.Run(port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
