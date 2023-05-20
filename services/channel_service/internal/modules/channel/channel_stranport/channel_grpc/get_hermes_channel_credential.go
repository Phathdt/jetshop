package channel_grpc

import (
	"context"

	"jetshop/common"
	jetshop_proto "jetshop/proto/out/proto"
	"jetshop/service-context/component/gormc"
	"jetshop/service-context/component/tracing"
	"jetshop/services/channel_service/internal/modules/channel/channel_repo"
	"jetshop/services/channel_service/internal/modules/channel/channel_storage"
)

func (s *channelGrpcServer) GetHermesChannelCredential(ctx context.Context, request *jetshop_proto.ChannelGetHermesCredentialRequest) (*jetshop_proto.ChannelGetHermesCredentialResponse, error) {
	ctx, span := tracing.WrapTraceIdFromIncomingContext(ctx, "grpc.get")
	defer span.End()

	sc := s.sc

	db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
	store := channel_storage.NewSqlStore(db)
	repo := channel_repo.NewRepo(store)

	cred, err := repo.GetChannelCredentialByCode(ctx, request.ChannelCode)
	if err != nil {
		panic(err)
	}

	return &jetshop_proto.ChannelGetHermesCredentialResponse{Cred: cred.ToProtoc()}, nil
}
