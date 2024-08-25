package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opa-oz/simple-queue/pkg"
	"github.com/opa-oz/simple-queue/pkg/utils"
)

const QueueHeader string = "X-From-Simple-Queue"

var (
	CannotGetTargets  = pkg.MessageResponse{Message: "Cannot get Targets"}
	TargetNotFound    = pkg.MessageResponse{Message: "Target not found"}
	CannotMarshalItem = pkg.MessageResponse{Message: "Cannot marshal json for the item"}
	NoQueueOpen       = pkg.MessageResponse{Message: "No queue open"}
)

var Yes = []string{"yes"}

func ScheduleGet(c *gin.Context) {
	target := c.Param("target")
	targets, err := utils.GetTargets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CannotGetTargets)
		return
	}

	incomingHeaders := c.Request.Header
	incomingPath := c.Param("request")
	incomingQuery := c.Request.URL.Query()

	incomingHeaders[QueueHeader] = Yes

	endpoint, ok := (*targets)[target]
	if !ok {
		c.JSON(http.StatusNotFound, TargetNotFound)
	}

	item := pkg.QueueItem{
		Header:   incomingHeaders,
		Path:     incomingPath,
		Query:    incomingQuery,
		Endpoint: endpoint,
		Method:   http.MethodGet,
	}

	qName := utils.GetQ(target)

	payload, err := item.MarshalBinary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CannotMarshalItem)
		return
	}

	queues, err := utils.GetRMQ(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.MessageResponse{Message: fmt.Sprintf("Failed to get queues: %e", err)})
		return
	}

	q, ok := (*queues)[qName]
	if !ok {
		c.JSON(http.StatusInternalServerError, NoQueueOpen)
		return
	}

	// https://github.com/adjust/rmq?tab=readme-ov-file
	err = (*q).PublishBytes(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.MessageResponse{Message: fmt.Sprintf("Failed to publish: %e", err)})
		return
	}

	c.String(http.StatusOK, string(payload))
}
