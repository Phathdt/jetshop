package chat_biz

import (
	"context"

	"go.opentelemetry.io/otel"
	"jetshop/integration/hermes"
	"jetshop/payloads"
	jetshop_proto "jetshop/proto/out/proto"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/component/watermillapp"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type PullDetailThreadRepo interface {
	UpsertThread(ctx context.Context, data []chat_model.Thread) error
}

type PullDetailThreadChannelRepo interface {
	GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto.HermesChannelCredential, error)
}

type pullDetailThreadBiz struct {
	repo        PullDetailThreadRepo
	channelRepo PullDetailThreadChannelRepo
	publisher   watermillapp.Publisher
	logger      sctx.Logger
}

func NewPullDetailThreadBiz(repo PullDetailThreadRepo, channelRepo PullDetailThreadChannelRepo, publisher watermillapp.Publisher, logger sctx.Logger) *pullDetailThreadBiz {
	return &pullDetailThreadBiz{repo: repo, channelRepo: channelRepo, publisher: publisher, logger: logger}
}

func (b *pullDetailThreadBiz) Response(ctx context.Context, channelCode, platformThreadId string) error {
	ctx, span := tracing.StartTrace(ctx, "biz.sync_thread")
	defer span.End()

	cred, err := b.channelRepo.GetHermesChannelCredentialByCode(ctx, channelCode)
	if err != nil {
		return err
	}

	client := hermes.NewClient()

	client.SetTracer(otel.Tracer("hermes"))

	thread, err := client.GetThread(ctx, cred.SellerId, platformThreadId)
	if err != nil {
		return err
	}

	threads := make([]chat_model.Thread, 1)
	threads[0] = chat_model.MapperToThread(thread)
	if err = b.repo.UpsertThread(ctx, threads); err != nil {
		return err
	}

	data := payloads.SyncMessageParams{
		ChannelCode:      cred.ChannelCode,
		PlatformThreadId: thread.ThreadId,
	}

	if err = b.publisher.Publish("sync_message", data); err != nil {
		b.logger.Errorln("publish message sync_message error = ", err)
	}

	return nil
}