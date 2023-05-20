package channel_storage

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"jetshop/service-context/component/tracing"
	"jetshop/services/channel_service/internal/modules/channel/channel_model"
)

type sqlStore struct {
	Db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{Db: db}
}

func (s *sqlStore) ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error) {
	ctx, span := tracing.WrapTraceIdFromIncomingContext(ctx, "sql_store.list")
	defer span.End()

	var credentials []channel_model.HermesChannelCredential

	db := s.Db.Table(channel_model.HermesChannelCredential{}.TableName()).Where(cond)

	if err := db.Select("*").
		Order("id desc").
		Find(&credentials).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return credentials, nil
}
