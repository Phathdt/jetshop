package chat_biz

import (
	"context"

	"jetshop/shared/proto/out/proto"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/tracing"
	"jetshop/shared/sctx/component/watermillapp"
	"jetshop/shared/sctx/core"
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
		if credential.SellerId == "" && credential.PlatformCode == "facebook" {
			if err = b.publisher.Publish("sync_thread", credential.ChannelCode); err != nil {
				b.logger.Errorln("publish message error = ", err)
			}
		}
	}

	return nil
}
