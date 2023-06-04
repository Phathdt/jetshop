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

func (s *sqlStore) GetThreadDetail(ctx context.Context, cond map[string]interface{}) (*chat_model.Thread, error) {
	ctx, span := tracing.StartTrace(ctx, "sql_store.get_thread_detail")
	defer span.End()

	var data chat_model.Thread

	db := s.db.WithContext(ctx).Table(chat_model.Thread{}.TableName()).Where(cond)

	result := db.Limit(1).Find(&data)

	if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errors.WithStack(errors.New("record not found"))
	}

	return &data, nil
}

func (s *sqlStore) ListThread(ctx context.Context, cond map[string]interface{}) ([]chat_model.Thread, error) {
	ctx, span := tracing.StartTrace(ctx, "sql_store.list")
	defer span.End()

	var data []chat_model.Thread

	db := s.db.WithContext(ctx).Table(chat_model.Thread{}.TableName())

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

func (s *sqlStore) UpsertThread(ctx context.Context, data []chat_model.Thread) error {
	ctx, span := tracing.StartTrace(ctx, "sql_store.upsert_thread")
	defer span.End()

	if err := s.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "channel_code"}, {Name: "platform_thread_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"customer_name",
			"customer_avatar_url",
			"updated_at",
		}),
	}).Create(&data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *sqlStore) UpdateThread(ctx context.Context, data []chat_model.Thread) error {
	ctx, span := tracing.StartTrace(ctx, "sql_store.update_thread")
	defer span.End()

	if err := s.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "channel_code"}, {Name: "platform_thread_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"from_type",
			"last_message",
			"send_time",
			"unread_count",
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

	if err := s.db.WithContext(ctx).Clauses(clause.OnConflict{
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

func (s *sqlStore) ListMessage(ctx context.Context, cond map[string]interface{}) ([]chat_model.Message, error) {
	ctx, span := tracing.StartTrace(ctx, "sql_store.list_message")
	defer span.End()

	var data []chat_model.Message

	db := s.db.WithContext(ctx).Table(chat_model.Message{}.TableName())

	if cursor, ok := cond["cursor"]; ok {
		db = db.Where("send_time < ?", cursor).Order("send_time desc")
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
