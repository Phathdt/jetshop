package chatcron

import (
	"context"
	"fmt"

	"jetshop/services/chat_service/internal/modules/chat_biz"
	"jetshop/shared/appgrpc"
	"jetshop/shared/common"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/tracing"
	"jetshop/shared/sctx/component/watermillapp"
)

func ScheduleSyncThread(sc sctx.ServiceContext) func() {
	return func() {
		ctx := context.Background()

		ctx, span := tracing.StartTrace(ctx, "cron.schedule_sync_thread")
		defer span.End()

		channelClient := sc.MustGet(common.KeyCompChannelClient).(appgrpc.ChannelClient)
		publisher := sc.MustGet(common.KeyCompNatsPub).(watermillapp.Publisher)
		logger := sctx.GlobalLogger().GetLogger("service")

		biz := chat_biz.NewScheduleSyncThreadBiz(channelClient, publisher, logger)

		if err := biz.Response(ctx); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
