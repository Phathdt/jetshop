package chat_storage

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"jetshop/service-context/component/tracing"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error) {
	ctx, span := tracing.StartTrace(ctx, "sql_store.list")
	defer span.End()

	var data []chat_model.Thread

	db := s.db.Table(chat_model.Thread{}.TableName())

	if cursor, ok := cond["cursor"]; ok {
		db = db.Where("last_message_time < ?", cursor).Order("last_message_time desc")
		delete(cond, "cursor")
	}

	if pageSize, ok := cond["page_size"]; ok {
		db = db.Limit(pageSize.(int))

		delete(cond, "page_size")
	}

	if err := db.Where(cond).Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return data, nil
}
