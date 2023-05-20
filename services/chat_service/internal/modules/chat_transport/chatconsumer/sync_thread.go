package chatconsumer

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"jetshop/appgrpc"
	"jetshop/common"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/tracing"
	"jetshop/services/chat_service/internal/modules/chat_biz"
)

func SyncThreadConsumer(sc sctx.ServiceContext) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		ctx := context.Background()

		ctx, span := tracing.StartTrace(ctx, "consumer.sync_thread")
		defer span.End()

		var channelCode string

		if err := json.Unmarshal(msg.Payload, &channelCode); err != nil {
			return err
		}

		channelClient := sc.MustGet(common.KeyCompChannelClient).(appgrpc.ChannelClient)

		biz := chat_biz.NewSyncThreadBiz(channelClient)

		if err := biz.Response(ctx, channelCode); err != nil {
			return err
		}

		return nil
	}
}
