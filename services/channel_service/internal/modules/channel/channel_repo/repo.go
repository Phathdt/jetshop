package channel_repo

import (
	"context"

	"jetshop/services/channel_service/internal/modules/channel/channel_model"
	"jetshop/shared/sctx/component/tracing"
)

type ChannelSqlStore interface {
	ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error)
	GetChannelCredentialByCondition(ctx context.Context, cond map[string]interface{}) (*channel_model.HermesChannelCredential, error)
}

type repo struct {
	store ChannelSqlStore
}

func NewRepo(store ChannelSqlStore) *repo {
	return &repo{store: store}
}

func (r *repo) ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error) {
	ctx, span := tracing.StartTrace(ctx, "repo.list")
	defer span.End()

	return r.store.ListChannelCredentials(ctx, cond)
}

func (r *repo) GetChannelCredentialByCode(ctx context.Context, channelCode string) (*channel_model.HermesChannelCredential, error) {
	ctx, span := tracing.StartTrace(ctx, "repo.get")
	defer span.End()

	return r.store.GetChannelCredentialByCondition(ctx, map[string]interface{}{"channel_code": channelCode})
}
