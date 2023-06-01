package chat_biz

import (
	"context"
	"time"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"jetshop/integration/hermes"
	"jetshop/integration/hermes/response"
	jetshop_proto "jetshop/proto/out/proto"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/component/watermillapp"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type syncThreadRepo interface {
	ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error)
}

type syncThreadChannelRepo interface {
	GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto.HermesChannelCredential, error)
}

type syncThreadBiz struct {
	repo        syncThreadRepo
	channelRepo syncThreadChannelRepo
	publisher   watermillapp.Publisher
	logger      sctx.Logger
}

func NewSyncThreadBiz(
	repo syncThreadRepo,
	channelRepo syncThreadChannelRepo,
	publisher watermillapp.Publisher,
	logger sctx.Logger) *syncThreadBiz {
	return &syncThreadBiz{
		repo:        repo,
		channelRepo: channelRepo,
		publisher:   publisher,
		logger:      logger,
	}
}

func (b *syncThreadBiz) Response(ctx context.Context, channelCode string) error {
	ctx, span := tracing.StartTrace(ctx, "biz.sync_thread")
	defer span.End()

	cred, err := b.channelRepo.GetHermesChannelCredentialByCode(ctx, channelCode)
	if err != nil {
		return err
	}

	client := hermes.NewClient()

	client.SetTracer(otel.Tracer("hermes"))

	t := time.Now()
	res, err := client.ListThread(ctx, cred.SellerId, t.UnixMilli(), 100)
	if err != nil {
		return err
	}

	threadIds := lo.Map(res.Data, func(datum response.Thread, index int) string {
		return datum.ThreadId
	})

	threads, err := b.repo.ListThread(ctx, map[string]interface{}{"channel_code": cred.ChannelCode, "platform_thread_id": threadIds})
	if err != nil {
		return err
	}

	threadMap := make(map[string]int64)
	lo.ForEach(threads, func(thread chat_model.Thread, index int) {
		threadMap[thread.PlatformThreadId] = thread.SendTime
	})

	var newThreadIds []string
	var needUpdateThreadIds []string

	lo.ForEach(res.Data, func(thread response.Thread, index int) {
		if threadMap[thread.ThreadId] == 0 {
			newThreadIds = append(newThreadIds, thread.ThreadId)
		}

		if threadMap[thread.ThreadId] != 0 && threadMap[thread.ThreadId] != thread.LastMessageTime {
			needUpdateThreadIds = append(needUpdateThreadIds, thread.ThreadId)
		}
	})

	if len(newThreadIds) != 0 {
		for _, newThreadId := range newThreadIds {
			data := make(map[string]string)
			data["channel_code"] = cred.ChannelCode
			data["platform_thread_id"] = newThreadId

			if err = b.publisher.Publish("detail_thread", data); err != nil {
				b.logger.Errorln("publish message detail_thread error = ", err)
			}
		}
	}

	if len(needUpdateThreadIds) != 0 {
		for _, needUpdateThreadId := range needUpdateThreadIds {
			data := make(map[string]string)
			data["channel_code"] = cred.ChannelCode
			data["platform_thread_id"] = needUpdateThreadId

			if err = b.publisher.Publish("sync_message", data); err != nil {
				b.logger.Errorln("publish message sync_message error = ", err)
			}
		}
	}

	return nil
}
