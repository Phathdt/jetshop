package chat_repo

import (
	"context"

	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/sctx/component/tracing"
)

type ChatStorage interface {
	GetThreadDetail(ctx context.Context, cond map[string]interface{}) (*chat_model.Thread, error)
	ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error)
	UpsertThread(ctx context.Context, data []chat_model.Thread) error
	UpsertMessage(ctx context.Context, data []chat_model.Message) error
}

type repo struct {
	store ChatStorage
}

func NewRepo(store ChatStorage) *repo {
	return &repo{store: store}
}

func (r *repo) ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error) {
	ctx, span := tracing.StartTrace(ctx, "repo.list")
	defer span.End()

	return r.store.ListThread(ctx, cond)
}

func (r *repo) UpsertThread(ctx context.Context, data []chat_model.Thread) error {
	ctx, span := tracing.StartTrace(ctx, "repo.upsert")
	defer span.End()

	return r.store.UpsertThread(ctx, data)
}

func (r *repo) UpsertMessages(ctx context.Context, data []chat_model.Message) error {
	ctx, span := tracing.StartTrace(ctx, "repo.upsert")
	defer span.End()

	return r.store.UpsertMessage(ctx, data)
}

func (r *repo) GetThreadDetail(ctx context.Context, cond map[string]interface{}) (*chat_model.Thread, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) ListMessage(ctx context.Context, cond map[string]interface{}) ([]chat_model.Message, error) {
	//TODO implement me
	panic("implement me")
}
