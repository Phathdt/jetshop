package postgresql

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"jetshop/pkg/service-context/component/tracing"
	"jetshop/pkg/service-context/core"
	"jetshop/services/product_service/internal/modules/productmodel"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r *repo) GetProduct(ctx context.Context, id int) (*productmodel.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "repo.get")
	defer span.End()

	var product productmodel.Product

	if err := r.db.
		WithContext(ctx).
		Table(product.TableName()).
		Where("id = ?", id).
		First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &product, nil
}
