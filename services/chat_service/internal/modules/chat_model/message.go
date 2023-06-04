package chat_model

import (
	"jetshop/services/chat_service/internal/modules/chat_enums"
	"jetshop/shared/common/enums"
	"jetshop/shared/sctx/core"
)

type Message struct {
	core.SQLModel
	ChannelCode       string                   `json:"channel_code"`
	Content           string                   `json:"content"`
	IsAutoReply       bool                     `json:"is_auto_reply"`
	MessageType       chat_enums.MessageType   `json:"message_type"`
	PlatformCode      enums.PlatformCode       `json:"platform_code"`
	PlatformMessageId string                   `json:"platform_message_id"`
	PlatformThreadId  string                   `json:"platform_thread_id"`
	SendTime          int64                    `json:"send_time"`
	FromType          chat_enums.FromType      `json:"from_type"`
	Status            chat_enums.MessageStatus `json:"status"`
	SentByUserId      int                      `json:"sent_by_user_id"`
	ChatbotProcessed  bool                     `json:"chatbot_processed"`
	AnswerMessageId   string                   `json:"answer_message_id"`
	AgentRequest      bool                     `json:"agent_request"`
	AutoReplyId       int                      `json:"auto_reply_id"`
}

func (m Message) TableName() string {
	return "messages"
}

func (m Message) ToClient() ClientMessage {
	var messageType chat_enums.ClientMessageType

	if m.AgentRequest {
		messageType = chat_enums.ClientMessageTypeSystem
	} else {
		switch m.MessageType {
		case chat_enums.MessageTypeSystem:
			messageType = chat_enums.ClientMessageTypeSystem
		case chat_enums.MessageTypeImage:
			messageType = chat_enums.ClientMessageTypeImage
		case chat_enums.MessageTypeSticker:
			messageType = chat_enums.ClientMessageTypeEmoji
		case chat_enums.MessageTypeProduct:
			messageType = chat_enums.ClientMessageTypeProduct
		case chat_enums.MessageTypeProductList:
			messageType = chat_enums.ClientMessageTypeProductList
		case chat_enums.MessageTypeOrder:
			messageType = chat_enums.ClientMessageTypeOrder
		case chat_enums.MessageTypeVideo:
			messageType = chat_enums.ClientMessageTypeVideo
		case chat_enums.MessageTypeVoucher:
			messageType = chat_enums.ClientMessageTypeVoucher
		case chat_enums.MessageTypeGeneric:
			messageType = chat_enums.ClientMessageTypeGeneric
		case chat_enums.MessageTypeAds:
			messageType = chat_enums.ClientMessageTypeAds
		case chat_enums.MessageTypeCarousel:
			messageType = chat_enums.ClientMessageTypeCarousel
		case chat_enums.MessageTypeTrigger:
			messageType = chat_enums.ClientMessageTypeSystem
		default:
			messageType = chat_enums.ClientMessageTypeText
		}
	}

	var status chat_enums.ClientMessageStatus
	switch m.Status {
	case chat_enums.MessageStatusReceived:
		status = chat_enums.ClientMessageStatusEdited
	case chat_enums.MessageStatusDeleted:
		status = chat_enums.ClientMessageStatusDeleted
	default:
		status = chat_enums.ClientMessageStatusSent
	}
	return ClientMessage{
		Id:                m.Id,
		SendTime:          m.SendTime,
		SentByUserId:      m.SentByUserId,
		Content:           m.Content,
		PlatformMessageId: m.PlatformMessageId,
		MessageType:       messageType,
		FromType:          m.FromType,
		Status:            status,
		PlatformId:        m.PlatformMessageId,
		Platform:          m.PlatformCode,
	}
}
