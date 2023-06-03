package appgrpc

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	jetshop_proto2 "jetshop/shared/proto/out/proto"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/common"
	"jetshop/shared/sctx/component/tracing"
)

type ChannelClient interface {
	ListHermesChannelCredential(ctx context.Context, isEnabled bool) ([]*jetshop_proto2.HermesChannelCredential, error)
	GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto2.HermesChannelCredential, error)
}

type channelClient struct {
	id         string
	consulHost string
	url        string
	logger     sctx.Logger
	client     jetshop_proto2.ChannelServiceClient
}

func NewChannelClient(id string) *channelClient {
	return &channelClient{id: id}
}

func (c *channelClient) ID() string {
	return c.id
}

func (c *channelClient) InitFlags() {
	c.consulHost = common.ConsulHost
}

func (c *channelClient) Activate(sc sctx.ServiceContext) error {
	c.logger = sc.Logger(c.id)

	c.logger.Infoln("Setup channel client service:", c.id)

	target := fmt.Sprintf("consul://%s/%s?healthy=true", c.consulHost, "channel_service")
	conn, err := grpc.Dial(
		target,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(500*time.Millisecond),
			otelgrpc.UnaryClientInterceptor()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		grpc.WithChainStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return err
	}

	c.client = jetshop_proto2.NewChannelServiceClient(conn)

	c.logger.Infof("Setup channel client service success with url %s", target)

	return nil
}

func (c *channelClient) Stop() error {
	c.logger.Infoln("channelClient grpc service stopped")

	return nil
}

func (c *channelClient) ListHermesChannelCredential(ctx context.Context, isEnabled bool) ([]*jetshop_proto2.HermesChannelCredential, error) {
	ctx, span := tracing.AppendTraceIdToOutgoingContext(ctx, "channel-client.get-list")
	defer span.End()

	rs, err := c.client.ListHermesChannelCredential(ctx, &jetshop_proto2.ChannelListHermesCredentialRequest{IsEnabled: isEnabled})
	if err != nil {
		return nil, err
	}

	return rs.Creds, nil
}

func (c *channelClient) GetHermesChannelCredentialByCode(ctx context.Context, channelCode string) (*jetshop_proto2.HermesChannelCredential, error) {
	ctx, span := tracing.AppendTraceIdToOutgoingContext(ctx, "channel-client.get-by-code")
	defer span.End()

	rs, err := c.client.GetHermesChannelCredential(ctx, &jetshop_proto2.ChannelGetHermesCredentialRequest{ChannelCode: channelCode})
	if err != nil {
		return nil, err
	}

	return rs.Cred, nil
}
