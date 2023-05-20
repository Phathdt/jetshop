package chat_biz

import (
	"context"
	"fmt"

	jetshop_proto "jetshop/proto/out/proto"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/core"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type scheduleSyncThreadChannelRepo interface {
	ListHermesChannelCredential(ctx context.Context, isEnabled bool) ([]*jetshop_proto.HermesChannelCredential, error)
}

type scheduleSyncThreadBiz struct {
	channelRepo scheduleSyncThreadChannelRepo
}

func NewScheduleSyncThreadBiz(channelRepo scheduleSyncThreadChannelRepo) *scheduleSyncThreadBiz {
	return &scheduleSyncThreadBiz{channelRepo: channelRepo}
}

func (b scheduleSyncThreadBiz) Response(ctx context.Context) error {
	ctx, span := tracing.StartTrace(ctx, "biz.schedule_sync_thread")
	defer span.End()

	credentials, err := b.channelRepo.ListHermesChannelCredential(ctx, true)

	if err != nil {
		return core.ErrInternalServerError.WithError(chat_model.ErrCannotListHermesChannelCredential.Error())
	}

	for _, credential := range credentials {
		fmt.Println(credential.ChannelCode)
	}

	return nil
}
