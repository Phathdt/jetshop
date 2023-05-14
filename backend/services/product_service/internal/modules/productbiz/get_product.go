package productbiz

import (
	"context"

	"jetshop/pkg/service-context/component/tracing"
	"jetshop/services/product_service/internal/modules/productmodel"
)

type GetProductRepo interface {
	GetProduct(ctx context.Context, id int) (*productmodel.Product, error)
}

type getProductBiz struct {
	repo GetProductRepo
}

func NewGetProductBiz(repo GetProductRepo) *getProductBiz {
	return &getProductBiz{repo: repo}
}

func (b *getProductBiz) Response(ctx context.Context, id int) (*productmodel.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "biz.get")
	defer span.End()

	return b.repo.GetProduct(ctx, id)
}
