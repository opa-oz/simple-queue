package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg/checks"
	"github.com/opa-oz/simple-queue/pkg/utils"
)

// @BasePath /api

// Ready godoc
// @Summary ready
// @Schemes
// @Description Check readiness
// @Tags utils
// @Accept json
// @Produce json
// @Success 200 {object} api.Ready.response
// @Router /ready [get]
func Ready(c *gin.Context) {
	type response struct {
		Message string `json:"message" swaggertype:"string" example:"Ready"`
	}

	rdb, err := utils.GetRedis(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response{Message: "Cannot get Redis"})
		return
	}

	err = checks.CheckRedis(c, rdb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{Message: "Redis is not working"})
		return
	}
	c.JSON(http.StatusOK, response{Message: "Ready"})
}
