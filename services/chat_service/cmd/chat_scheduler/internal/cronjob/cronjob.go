package cronjob

import (
	"jetshop/services/chat_service/internal/modules/chat_transport/chatcron"
	"jetshop/shared/common"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/cronjob"
)

func NewCronjob(sc sctx.ServiceContext) {
	c := sc.MustGet(common.KeyCron).(cronjob.Cronjob)

	c.Enqueue("* * * * * *", chatcron.ScheduleSyncThread(sc))
}
