package channel_grpc

import (
	"context"

	"jetshop/common"
	jetshop_proto "jetshop/proto/out/proto"
	"jetshop/service-context/component/gormc"
	"jetshop/services/channel_service/internal/modules/channel/channel_repo"
	"jetshop/services/channel_service/internal/modules/channel/channel_storage"
)

func (s *channelGrpcServer) ListHermesChannelCredential(ctx context.Context, request *jetshop_proto.ChannelListHermesCredentialRequest) (*jetshop_proto.ChannelListHermesCredentialResponse, error) {
	sc := s.sc

	db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
	store := channel_storage.NewSqlStore(db)
	repo := channel_repo.NewRepo(store)

	credentials, err := repo.ListChannelCredentials(ctx, map[string]interface{}{"is_enabled": request.IsEnabled})
	if err != nil {
		panic(err)
	}

	res := make([]*jetshop_proto.HermesChannelCredential, len(credentials))

	for i, credential := range credentials {
		res[i] = credential.ToProtoc()
	}

	return &jetshop_proto.ChannelListHermesCredentialResponse{Creds: res}, nil
}
