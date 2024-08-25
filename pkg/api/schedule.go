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

	qName := utils.GetQ(target)

	payload, err := item.MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "Cannot marshal json for the item"})
		return
	}

	queues, err := utils.GetRMQ(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: fmt.Sprintf("Failed to get queues: %e", err)})
		return
	}

	q, ok := (*queues)[qName]
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: "No queue open"})
		return
	}

	// https://github.com/adjust/rmq?tab=readme-ov-file
	err = (*q).PublishBytes(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.MessageResponse{Message: fmt.Sprintf("Failed to publish: %e", err)})
		return
	}

	c.String(http.StatusOK, string(payload))
}
