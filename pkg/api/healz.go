package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/opa-oz/simple-queue/pkg/checks"
	"github.com/opa-oz/simple-queue/pkg/utils"
)

var (
	CannotGetRedis    = pkg.MessageResponse{Message: "Cannot get Redis"}
	RedisIsNotWorking = pkg.MessageResponse{Message: "Redis is not working"}
	OK                = pkg.MessageResponse{Message: "OK"}
)

// @BasePath /api

// Healz godoc
// @Summary healz
// @Schemes
// @Description Check health endpoint
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} utils.MessageResponse
// @Router /healz [get]
func Healz(c *gin.Context) {
	rdb, err := utils.GetRedis(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, CannotGetRedis)
		return
	}

	err = checks.CheckRedis(c, rdb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RedisIsNotWorking)
		return
	}

	c.JSON(http.StatusOK, OK)
}
