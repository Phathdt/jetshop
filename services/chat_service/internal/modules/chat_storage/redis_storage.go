package chat_storage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"jetshop/services/chat_service/internal/modules/chat_model"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}

func (r redisStore) PublishMessage(ctx context.Context, channelCode, threadId string, messages []chat_model.Message) error {
	data := map[string]interface{}{
		"kind":         "messages",
		"channel_code": channelCode,
		"data": map[string]interface{}{
			"conversation_id": threadId,
			"messages": lo.Map(messages, func(m chat_model.Message, index int) chat_model.ClientMessage {
				return m.ToClient()
			}),
		},
	}

	if err := r.client.Publish(ctx, "chat-broadcast", data).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
