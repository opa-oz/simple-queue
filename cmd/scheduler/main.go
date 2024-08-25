package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg/api"
	"github.com/opa-oz/simple-queue/pkg/config"
	"github.com/opa-oz/simple-queue/pkg/middlewares"
	"github.com/opa-oz/simple-queue/pkg/redis"
	"github.com/opa-oz/simple-queue/pkg/utils"
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

	errChan := make(chan error, 10)
	go utils.LogErrors(errChan)

	connection, err := redis.GetRMQConnection(rdb, errChan)
	if err != nil {
		fmt.Println(err)
		return
	}

	targets, err := config.GetTargets(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	queues, err := config.PrepareQueues(connection, targets, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Use(middlewares.RequestLogger())
	r.Use(middlewares.ResponseLogger())
	r.Use(middlewares.RedisMiddleware(rdb))
	r.Use(middlewares.RMQMiddleware(queues))
	r.Use(middlewares.CfgMiddleware(cfg))

	r.GET("/healz", api.Healz)
	r.GET("/ready", api.Ready)

	rg := r.Group("/simple")
	{
		rg.GET("/:target/*request", api.ScheduleGet)
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	err = r.Run(port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
