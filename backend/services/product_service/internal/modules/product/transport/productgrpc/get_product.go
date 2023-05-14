package productgrpc

import (
	"context"

	"jetshop/common"
	"jetshop/pkg/service-context/component/gormc"
	"jetshop/pkg/service-context/component/redisc"
	"jetshop/pkg/service-context/component/tracing"
	protos "jetshop/proto/out/proto"
	"jetshop/services/product_service/internal/modules/product/repository"
	"jetshop/services/product_service/internal/modules/product/storage"
)

func (s *productGrpcServer) GetProduct(ctx context.Context, request *protos.GetProductRequest) (*protos.GetProductResponse, error) {
	ctx, span := tracing.WrapTraceIdFromIncomingContext(ctx, "grpc.get")
	defer span.End()

	db := s.sc.MustGet(common.KeyCompGorm).(gormc.GormComponent)
	rdClient := s.sc.MustGet(common.KeyCompRedis).(redisc.RedisComponent)

	store := storage.NewSqlStore(db.GetDB())
	redisStore := storage.NewRedisStore(rdClient.GetClient())
	repo := repository.NewRepo(redisStore, store)

	product, err := repo.GetProduct(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &protos.GetProductResponse{Product: &protos.Product{
		Id:    uint32(product.Id),
		Name:  product.Name,
		Price: uint32(product.Price),
	}}, nil
}
