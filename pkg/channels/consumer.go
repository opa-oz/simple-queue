package channels

import (
	"fmt"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/opa-oz/simple-queue/pkg"
)

var (
	consumeDuration = time.Millisecond
	reportBatchSize = 10000
)

type Consumer struct {
	name   string
	count  int
	before time.Time
}

func NewConsumer() *Consumer {
	return &Consumer{
		name:   "consumer",
		count:  0,
		before: time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	payload := delivery.Payload()

	item := pkg.QueueItem{}
	err := item.UnmarshalBinary([]byte(payload))

	if err != nil {
		// handle json error
		if err := delivery.Reject(); err != nil {
			// handle reject error
		}
		return
	}

	fmt.Println("Got payload for", item.Endpoint, item.Path)
}
