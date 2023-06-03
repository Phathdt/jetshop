package channel_grpc

import (
	"jetshop/shared/sctx"
)

type channelGrpcServer struct {
	sc sctx.ServiceContext
}

func NewChannelGrpcServer(sc sctx.ServiceContext) *channelGrpcServer {
	return &channelGrpcServer{sc: sc}
}
