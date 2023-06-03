package chat_storage

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/sctx/component/tracing"
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

func (s *sqlStore) UpsertConversation(ctx context.Context, data []chat_model.Thread) error {
	ctx, span := tracing.StartTrace(ctx, "sql_store.upsert_thread")
	defer span.End()

	if err := s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "channel_code"}, {Name: "platform_thread_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"customer_name",
			"customer_avatar_url",
			"unread_count",
			"last_message",
			"send_time",
			"from_type",
			"last_message_is_auto_reply",
			"bot_stop_at",
			"op_source",
			"op_source_send_time",
			"updated_at",
		}),
	}).Create(&data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *sqlStore) UpsertMessage(ctx context.Context, data []chat_model.Message) error {
	ctx, span := tracing.StartTrace(ctx, "sql_store.upsert_message")
	defer span.End()

	if err := s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "platform_thread_id"}, {Name: "platform_message_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"send_time",
			"content",
			"status",
			"updated_at",
		}),
	}).Create(&data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
