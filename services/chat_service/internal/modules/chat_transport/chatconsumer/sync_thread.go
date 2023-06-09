package chatconsumer

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"jetshop/services/chat_service/internal/modules/chat_biz"
	"jetshop/services/chat_service/internal/modules/chat_repo"
	"jetshop/services/chat_service/internal/modules/chat_storage"
	"jetshop/shared/appgrpc"
	"jetshop/shared/common"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/gormc"
	"jetshop/shared/sctx/component/tracing"
	"jetshop/shared/sctx/component/watermillapp"
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
		publisher := sc.MustGet(common.KeyCompNatsPub).(watermillapp.Publisher)
		logger := sctx.GlobalLogger().GetLogger("service")

		db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
		store := chat_storage.NewSqlStore(db)
		repo := chat_repo.NewRepo(store)

		biz := chat_biz.NewSyncThreadBiz(repo, channelClient, publisher, logger)

		if err := biz.Response(ctx, channelCode); err != nil {
			return err
		}

		return nil
	}
}
