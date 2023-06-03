package chatconsumer

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"jetshop/appgrpc"
	"jetshop/common"
	"jetshop/payloads"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/gormc"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/component/watermillapp"
	"jetshop/services/chat_service/internal/modules/chat_biz"
	"jetshop/services/chat_service/internal/modules/chat_repo"
	"jetshop/services/chat_service/internal/modules/chat_storage"
)

func SyncMessageConsumer(sc sctx.ServiceContext) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		ctx := context.Background()

		ctx, span := tracing.StartTrace(ctx, "consumer.detail_thread")
		defer span.End()

		var payload payloads.PullDetailThreadParams

		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return err
		}

		channelClient := sc.MustGet(common.KeyCompChannelClient).(appgrpc.ChannelClient)
		publisher := sc.MustGet(common.KeyCompNatsPub).(watermillapp.Publisher)
		logger := sctx.GlobalLogger().GetLogger("service")

		db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
		store := chat_storage.NewSqlStore(db)
		repo := chat_repo.NewRepo(store)

		biz := chat_biz.NewSyncMessageBiz(repo, channelClient, publisher, logger)

		if err := biz.Response(ctx, payload.ChannelCode, payload.PlatformThreadId); err != nil {
			return err
		}

		return nil
	}
}
