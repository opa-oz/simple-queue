package consumers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/opa-oz/go-todo/todo"
	"github.com/opa-oz/simple-queue/pkg"
)

type HttpConsumer struct {
	name   string
	client *http.Client
	before time.Time
}

func NewHttpConsumer() *HttpConsumer {
	return &HttpConsumer{
		name:   todo.String("Proper name for Consumer", "consumer"),
		client: &http.Client{},
		before: time.Now(),
	}
}

func (consumer *HttpConsumer) Consume(delivery rmq.Delivery) {
	payload := delivery.Payload()

	item := pkg.QueueItem{}
	err := item.UnmarshalBinary([]byte(payload))

	if err != nil {
		fmt.Println(todo.String("Replace with proper logging", err.Error()))
		delivery.Reject()
		return
	}

	endpoint := fmt.Sprintf("%s/%s", strings.TrimRight(item.Endpoint, pkg.Slash), strings.TrimLeft(item.Path, pkg.Slash))
	q := item.Query.Encode()

	uri := fmt.Sprintf("%s?%s", endpoint, q)
	r, err := http.NewRequest(item.Method, uri, nil)
	if err != nil {
		fmt.Println(todo.String("Replace with proper logging", err.Error()))
		delivery.Reject()
		return
	}

	fmt.Println(todo.String("Debug log", uri))

	r.Header = item.Header
	res, err := consumer.client.Do(r)
	if err != nil {
		fmt.Println(todo.String("Replace with proper logging", err.Error()))
		delivery.Reject()
		return
	}

	if res.StatusCode > 400 {
		delivery.Reject()
	} else {
		delivery.Ack()
	}
}
