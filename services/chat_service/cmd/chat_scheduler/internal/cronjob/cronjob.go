package cronjob

import (
	"jetshop/common"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/cronjob"
	"jetshop/services/chat_service/internal/modules/chat_transport/chatcron"
)

func NewCronjob(sc sctx.ServiceContext) {
	c := sc.MustGet(common.KeyCron).(cronjob.Cronjob)

	c.Enqueue("* * * * * *", chatcron.ScheduleSyncThread(sc))
}
