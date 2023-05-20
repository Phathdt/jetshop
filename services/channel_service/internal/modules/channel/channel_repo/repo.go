package channel_repo

import (
	"context"

	"jetshop/service-context/component/tracing"
	"jetshop/services/channel_service/internal/modules/channel/channel_model"
)

type ChannelSqlStore interface {
	ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error)
}

type repo struct {
	store ChannelSqlStore
}

func NewRepo(store ChannelSqlStore) *repo {
	return &repo{store: store}
}

func (r *repo) ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error) {
	ctx, span := tracing.WrapTraceIdFromIncomingContext(ctx, "repo.list")
	defer span.End()

	return r.store.ListChannelCredentials(ctx, cond)
}
