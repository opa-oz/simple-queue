package channels

import (
	"context"
	"fmt"

	"github.com/opa-oz/simple-queue/pkg/utils"
	"github.com/redis/go-redis/v9"
)

func GoToThread(rdb *redis.Client, target string) error {
	ctx := context.TODO()
	pubsub := rdb.Subscribe(ctx, utils.GetQ(target))

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}
}
