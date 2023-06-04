package chat_biz

import (
	"context"
	"errors"
	"time"

	"github.com/samber/lo"
	"jetshop/services/chat_service/internal/modules/chat_enums"
	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/sctx/component/tracing"
)

type UpdateThreadRepo interface {
	GetThreadDetail(ctx context.Context, cond map[string]interface{}) (*chat_model.Thread, error)
	UpsertThread(ctx context.Context, data []chat_model.Thread) error
	ListMessage(ctx context.Context, cond map[string]interface{}) ([]chat_model.Message, error)
}

type updateThreadBiz struct {
	repo UpdateThreadRepo
}

func NewUpdateThreadBiz(repo UpdateThreadRepo) *updateThreadBiz {
	return &updateThreadBiz{repo: repo}
}

func (b *updateThreadBiz) Response(ctx context.Context, channelCode, platformThreadId string) error {
	ctx, span := tracing.StartTrace(ctx, "biz.update_thread")
	defer span.End()

	thread, err := b.repo.GetThreadDetail(ctx, map[string]interface{}{"channel_code": channelCode, "platform_thread_id": platformThreadId})
	if err != nil {
		return err
	}

	messages, err := b.repo.ListMessage(ctx, map[string]interface{}{"channel_code": channelCode, "platform_thread_id": platformThreadId})
	if err != nil {
		return err
	}

	latestMessage, ok := lo.Find(messages, func(m chat_model.Message) bool {
		return m.MessageType != chat_enums.MessageTypeSystem && m.MessageType != chat_enums.MessageTypeTrigger
	})
	if !ok {
		return errors.New("cannot find latest message")
	}

	b.initAttrs(thread, latestMessage, messages)
	b.calculateUnreadCount(thread, latestMessage, messages)

	t := time.Now()
	thread.UpdatedAt = &t

	if err = b.repo.UpsertThread(ctx, []chat_model.Thread{*thread}); err != nil {
		return err
	}

	return nil
}

func (b *updateThreadBiz) initAttrs(thread *chat_model.Thread, latestMessage chat_model.Message, messages []chat_model.Message) {
	switch latestMessage.FromType {
	case chat_enums.FromTypeBuyer:
		thread.FromType = chat_enums.FromTypeBuyer
		thread.LastMessageIsAutoReply = false

	case chat_enums.FromTypeSeller:
		if latestMessage.IsAutoReply {
			thread.FromType = chat_enums.FromTypeBuyer
			thread.LastMessageIsAutoReply = false
		} else {
			thread.FromType = chat_enums.FromTypeSeller
			thread.LastMessageIsAutoReply = false
		}

	case chat_enums.FromTypeAuto:
		thread.FromType = chat_enums.FromTypeBuyer
		thread.LastMessageIsAutoReply = true

	default:
		thread.FromType = chat_enums.FromTypeSeller
		thread.LastMessageIsAutoReply = false
	}
}

func (b *updateThreadBiz) calculateUnreadCount(thread *chat_model.Thread, latestMessage chat_model.Message, messages []chat_model.Message) {
	switch latestMessage.FromType {
	case chat_enums.FromTypeSeller:
		if latestMessage.IsAutoReply {
			thread.UnreadCount = 1
		} else {
			thread.UnreadCount = 0
		}
	case chat_enums.FromTypeAuto:
		thread.UnreadCount = 1
	default:

		latestSellerMessage, ok := lo.Find(messages, func(m chat_model.Message) bool {
			return m.FromType == chat_enums.FromTypeSeller
		})

		if ok {
			sellerMessages := lo.Filter(messages, func(m chat_model.Message, index int) bool {
				return m.FromType == chat_enums.FromTypeBuyer && m.SendTime > latestSellerMessage.SendTime
			})

			thread.UnreadCount = len(sellerMessages)
		} else {
			sellerMessages := lo.Filter(messages, func(m chat_model.Message, index int) bool {
				return m.FromType == chat_enums.FromTypeBuyer
			})

			thread.UnreadCount = len(sellerMessages)
		}
	}

}
