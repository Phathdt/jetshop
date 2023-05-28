package chat_repo

import (
	"context"

	"jetshop/service-context/component/tracing"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type ChatStorage interface {
	ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error)
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
