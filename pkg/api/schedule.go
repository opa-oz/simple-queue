package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/opa-oz/simple-queue/pkg/utils"
)

const QueueHeader string = "X-From-Simple-Queue"

var Yes = []string{"yes"}

func ScheduleGet(c *gin.Context) {
	rdb, err := utils.GetRedis(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "Cannot get Redis"})
		return
	}

	target := c.Param("target")

	incomingHeaders := c.Request.Header
	incomingPath := c.Param("request")
	incomingQuery := c.Request.URL.Query()

	incomingHeaders[QueueHeader] = Yes

	item := pkg.QueueItem{
		Header:   incomingHeaders,
		Path:     incomingPath,
		Query:    incomingQuery,
		Endpoint: todo.String("Replace with target map", target),
	}

	payload, err := item.MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "Cannot marshal json for the item"})
		return
	}

	// https://dev.to/franciscomendes10866/using-redis-pub-sub-with-golang-mf9
	err = rdb.Publish(c, utils.GetQ(target), payload).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: fmt.Sprintf("Failed to publish: %e", err)})
		return
	}

	c.String(http.StatusOK, string(payload))
}
