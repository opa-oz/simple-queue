package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/opa-oz/simple-queue/pkg/checks"
	"github.com/opa-oz/simple-queue/pkg/utils"
)

var (
	IAmReady = pkg.MessageResponse{Message: "Ready"}
)

// @BasePath /api

// Ready godoc
// @Summary ready
// @Schemes
// @Description Check readiness
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} utils.MessageResponse
// @Router /ready [get]
func Ready(c *gin.Context) {
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
	c.JSON(http.StatusOK, IAmReady)
}
