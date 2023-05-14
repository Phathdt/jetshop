package repository

import (
	"context"

	"jetshop/pkg/service-context/component/tracing"
	"jetshop/services/product_service/internal/modules/product/model"
)

type ProductStorage interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
}

type repo struct {
	storage ProductStorage
}

func NewRepo(storage ProductStorage) *repo {
	return &repo{storage: storage}
}

func (r *repo) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "repository.get")
	defer span.End()

	return r.storage.GetProduct(ctx, id)
}
