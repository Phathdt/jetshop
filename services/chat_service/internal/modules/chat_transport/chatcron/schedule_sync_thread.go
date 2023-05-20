package chatcron

import (
	"context"
	"fmt"

	"jetshop/appgrpc"
	"jetshop/common"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/tracing"
	"jetshop/services/chat_service/internal/modules/chat_biz"
)

func ScheduleSyncThread(sc sctx.ServiceContext) func() {
	return func() {
		ctx := context.Background()

		ctx, span := tracing.StartTrace(ctx, "cron.schedule_sync_thread")
		defer span.End()

		client := sc.MustGet(common.KeyCompChannelClient).(appgrpc.ChannelClient)

		biz := chat_biz.NewScheduleSyncThreadBiz(client)

		if err := biz.Response(ctx); err != nil {
			fmt.Println("111111111")
			fmt.Println(err)
			panic(err)
		}
	}
}
