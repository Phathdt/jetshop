package chat_biz

import (
	"context"
	"time"

	"jetshop/services/chat_service/internal/modules/chat_mapper"
	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/integration/hermes"
	"jetshop/shared/payloads"
	"jetshop/shared/proto/out/proto"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/tracing"
	"jetshop/shared/sctx/component/watermillapp"
)

type SyncMessageRepo interface {
	UpsertMessages(ctx context.Context, data []chat_model.Message) error
}

type SyncMessageChannelRepo interface {
	GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto.HermesChannelCredential, error)
}

type syncMessageBiz struct {
	repo        SyncMessageRepo
	channelRepo SyncMessageChannelRepo
	publisher   watermillapp.Publisher
	logger      sctx.Logger
}

func NewSyncMessageBiz(repo SyncMessageRepo, channelRepo SyncMessageChannelRepo, publisher watermillapp.Publisher, logger sctx.Logger) *syncMessageBiz {
	return &syncMessageBiz{repo: repo, channelRepo: channelRepo, publisher: publisher, logger: logger}
}

func (b *syncMessageBiz) Response(ctx context.Context, channelCode, platformThreadId string) error {
	ctx, span := tracing.StartTrace(ctx, "biz.sync_thread")
	defer span.End()

	cred, err := b.channelRepo.GetHermesChannelCredentialByCode(ctx, channelCode)
	if err != nil {
		return err
	}

	client := hermes.NewClient()

	t := time.Now()
	rs, err := client.ListMessage(ctx, cred.SellerId, platformThreadId, t.UnixMilli(), 100)
	if err != nil {
		return err
	}

	messages := make([]chat_model.Message, len(rs.Data))
	for i, m := range rs.Data {
		message, err := chat_mapper.MapperToMessage(&m)
		if err != nil {
			return err
		}

		messages[i] = *message
	}

	if err = b.repo.UpsertMessages(ctx, messages); err != nil {
		return err
	}

	data := payloads.UpdateThreadParams{
		ChannelCode:      channelCode,
		PlatformThreadId: platformThreadId,
	}

	if err = b.publisher.Publish("update_thread", data); err != nil {
		b.logger.Errorln("publish message update_thread error = ", err)
	}

	return nil
}
