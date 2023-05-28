package chat_biz

import (
	"context"

	jetshop_proto "jetshop/proto/out/proto"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/component/watermillapp"
	"jetshop/service-context/core"
)

type scheduleSyncThreadChannelRepo interface {
	ListHermesChannelCredential(ctx context.Context, isEnabled bool) ([]*jetshop_proto.HermesChannelCredential, error)
}

type scheduleSyncThreadBiz struct {
	channelRepo scheduleSyncThreadChannelRepo
	publisher   watermillapp.Publisher
	logger      sctx.Logger
}

func NewScheduleSyncThreadBiz(channelRepo scheduleSyncThreadChannelRepo, publisher watermillapp.Publisher, logger sctx.Logger) *scheduleSyncThreadBiz {
	return &scheduleSyncThreadBiz{channelRepo: channelRepo, publisher: publisher, logger: logger}
}

func (b scheduleSyncThreadBiz) Response(ctx context.Context) error {
	ctx, span := tracing.StartTrace(ctx, "biz.schedule_sync_thread")
	defer span.End()

	credentials, err := b.channelRepo.ListHermesChannelCredential(ctx, true)

	if err != nil {
		return core.ErrInternalServerError.WithError(err.Error())
	}

	for _, credential := range credentials {
		if err = b.publisher.Publish("sync_thread", credential.ChannelCode); err != nil {
			b.logger.Errorln("publish message error = ", err)
		}
	}

	return nil
}
