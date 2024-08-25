package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg/checks"
	"github.com/opa-oz/simple-queue/pkg/utils"
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
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "Cannot get Redis"})
		return
	}

	err = checks.CheckRedis(c, rdb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "Redis is not working"})
		return
	}

	c.JSON(http.StatusOK, utils.MessageResponse{Message: "OK"})
}
