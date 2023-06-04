package chatconsumer

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"jetshop/services/chat_service/internal/modules/chat_biz"
	"jetshop/services/chat_service/internal/modules/chat_repo"
	"jetshop/services/chat_service/internal/modules/chat_storage"
	"jetshop/shared/common"
	"jetshop/shared/payloads"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/gormc"
	"jetshop/shared/sctx/component/tracing"
)

func UpdateThread(sc sctx.ServiceContext) func(msg *message.Message) error {
	return func(msg *message.Message) error {
		ctx := context.Background()

		ctx, span := tracing.StartTrace(ctx, "consumer.detail_thread")
		defer span.End()

		var payload payloads.PullDetailThreadParams

		if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			return err
		}

		db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
		store := chat_storage.NewSqlStore(db)
		repo := chat_repo.NewRepo(store)

		biz := chat_biz.NewUpdateThreadBiz(repo)

		if err := biz.Response(ctx, payload.ChannelCode, payload.PlatformThreadId); err != nil {
			return err
		}

		return nil
	}
}