package business

import (
	"context"

	"jetshop/pkg/service-context/component/tracing"
	"jetshop/pkg/service-context/core"
	"jetshop/services/product_service/internal/modules/product/model"
)

type GetProductRepo interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
}

type getProductBiz struct {
	repo GetProductRepo
}

func NewGetProductBiz(repo GetProductRepo) *getProductBiz {
	return &getProductBiz{repo: repo}
}

func (b *getProductBiz) Response(ctx context.Context, id int) (*model.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "biz.get")
	defer span.End()

	product, err := b.repo.GetProduct(ctx, id)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetProduct.Error()).
			WithDebug(err.Error())
	}

	return product, nil
}