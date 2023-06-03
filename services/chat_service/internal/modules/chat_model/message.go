package chat_model

import (
	"jetshop/integration/hermes/response"
	"jetshop/service-context/core"
)

type Message struct {
	core.SQLModel
	ChannelCode       string `json:"channel_code"`
	Content           string `json:"content"`
	IsAutoReply       bool   `json:"is_auto_reply"`
	MessageType       string `json:"message_type"`
	PlatformCode      string `json:"platform_code"`
	PlatformMessageId string `json:"platform_message_id"`
	PlatformThreadId  string `json:"platform_thread_id"`
	SendTime          int64  `json:"send_time"`
	FromType          string `json:"from_type"`
	Status            string `json:"status"`
	SentByUserId      int    `json:"sent_by_user_id"`
	ChatbotProcessed  bool   `json:"chatbot_processed"`
	AnswerMessageId   string `json:"answer_message_id"`
	AgentRequest      bool   `json:"agent_request"`
	AutoReplyId       int    `json:"auto_reply_id"`
}

func (m Message) Message() string {
	return "messages"
}

func MapperToMessage(message *response.Message) Message {
	return Message{
		SQLModel:          *core.NewUpsertWithoutIdSQLModel(),
		ChannelCode:       message.ChannelCode,
		PlatformCode:      message.Platform,
		PlatformThreadId:  message.ThreadId,
		PlatformMessageId: message.MessageId,
		Content:           message.Content,
		//	message_type
		// fromtype
		SendTime: message.SendTime,
		//status
		IsAutoReply: message.AutoReply,
	}
}
