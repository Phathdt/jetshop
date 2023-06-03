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

func (m Message) Message() string {
	return "messages"
}
