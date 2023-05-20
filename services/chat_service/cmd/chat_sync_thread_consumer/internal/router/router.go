package router

import (
	"jetshop/common"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/watermillapp"
	"jetshop/services/chat_service/internal/modules/chat_transport/chatconsumer"
)

func NewRouter(sc sctx.ServiceContext) {
	c := sc.MustGet(common.KeyCompNatsSub).(watermillapp.Subscriber)

	c.AddNoPublisherHandler("sync_thread", "sync_thread", chatconsumer.SyncThreadConsumer(sc))
}
