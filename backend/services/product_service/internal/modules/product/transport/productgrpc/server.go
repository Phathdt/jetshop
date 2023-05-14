package productgrpc

import (
	sctx "jetshop/pkg/service-context"
)

type productGrpcServer struct {
	sc sctx.ServiceContext
}

func NewProductGrpcServer(sc sctx.ServiceContext) *productGrpcServer {
	return &productGrpcServer{sc: sc}
}
