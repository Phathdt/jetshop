package repository

import (
	"context"

	"jetshop/pkg/service-context/component/tracing"
	"jetshop/services/product_service/internal/modules/product/model"
)

type ProductStorage interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
}

type ProductCacheStore interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
	SetProduct(ctx context.Context, product *model.Product) error
}

type repo struct {
	cacheStore ProductCacheStore
	storage    ProductStorage
}

func NewRepo(cacheStore ProductCacheStore, storage ProductStorage) *repo {
	return &repo{cacheStore: cacheStore, storage: storage}
}

func (r *repo) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "repository.get")
	defer span.End()

	if product, _ := r.cacheStore.GetProduct(ctx, id); product != nil {
		return product, nil
	}

	product, err := r.storage.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = r.cacheStore.SetProduct(ctx, product)

	return product, nil
}
