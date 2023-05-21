package chat_biz

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"jetshop/integration/hermes"
	jetshop_proto "jetshop/proto/out/proto"
	"jetshop/service-context/component/tracing"
)

type syncThreadChannelRepo interface {
	GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto.HermesChannelCredential, error)
}

type syncThreadBiz struct {
	channelRepo syncThreadChannelRepo
}

func NewSyncThreadBiz(channelRepo syncThreadChannelRepo) *syncThreadBiz {
	return &syncThreadBiz{channelRepo: channelRepo}
}

func (b *syncThreadBiz) Response(ctx context.Context, channelCode string) error {
	ctx, span := tracing.StartTrace(ctx, "biz.sync_thread")
	defer span.End()

	cred, err := b.channelRepo.GetHermesChannelCredentialByCode(ctx, channelCode)
	if err != nil {
		return err
	}

	if cred.SellerId == "" || cred.PlatformCode != "facebook" {
		return nil
	}

	client := hermes.NewClient()

	client.SetTracer(otel.Tracer("hermes"))

	t := time.Now()
	res, err := client.ListThread(ctx, cred.SellerId, t.UnixMilli(), 100)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
