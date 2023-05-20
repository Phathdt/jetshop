package chatconsumer

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	sctx "jetshop/service-context"
)

func UpsertMessageConsumer(sc sctx.ServiceContext) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		var channelCode string

		if err := json.Unmarshal(msg.Payload, &channelCode); err != nil {
			return err
		}

		fmt.Println("channelCode", channelCode)
		return nil
	}
}
