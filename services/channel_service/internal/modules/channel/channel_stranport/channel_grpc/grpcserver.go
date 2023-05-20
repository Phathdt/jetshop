package channel_grpc

import (
	sctx "jetshop/service-context"
)

type channelGrpcServer struct {
	sc sctx.ServiceContext
}

func NewChannelGrpcServer(sc sctx.ServiceContext) *channelGrpcServer {
	return &channelGrpcServer{sc: sc}
}
