package storage

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/core"
	"jetshop/services/product_service/internal/modules/product/model"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (r *sqlStore) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "storage.get")
	defer span.End()

	var product model.Product

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
