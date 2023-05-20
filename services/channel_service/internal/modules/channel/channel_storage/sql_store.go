package channel_storage

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"jetshop/service-context/component/tracing"
	"jetshop/service-context/core"
	"jetshop/services/channel_service/internal/modules/channel/channel_model"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) ListChannelCredentials(ctx context.Context, cond map[string]interface{}) ([]channel_model.HermesChannelCredential, error) {
	ctx, span := tracing.StartTrace(ctx, "sql_store.list")
	defer span.End()

	var credentials []channel_model.HermesChannelCredential

	db := s.db.WithContext(ctx).Table(channel_model.HermesChannelCredential{}.TableName()).Where(cond)

	if err := db.Select("*").
		Order("id desc").
		Find(&credentials).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return credentials, nil
}

func (s *sqlStore) GetChannelCredentialByCondition(ctx context.Context, cond map[string]interface{}) (*channel_model.HermesChannelCredential, error) {
	var data channel_model.HermesChannelCredential

	db := s.db.WithContext(ctx).Table(channel_model.HermesChannelCredential{}.TableName())

	result := db.Where(cond).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, core.ErrRecordNotFound
	}

	return &data, nil
}
