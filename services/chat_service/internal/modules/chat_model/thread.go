package chat_model

import (
	"jetshop/integration/hermes/response"
	"jetshop/service-context/core"
)

type Thread struct {
	core.SQLModel
	ChannelCode            string `json:"channel_code"`
	PlatformThreadId       string `json:"platform_thread_id"`
	PlatformCustomerId     string `json:"platform_customer_id"`
	CustomerName           string `json:"customer_name"`
	EncodedCustomerName    string `json:"encoded_customer_name"`
	CustomerAvatarUrl      string `json:"customer_avatar_url"`
	UnreadCount            int    `json:"unread_count"`
	PlatformCode           string `json:"platform_code"`
	LastMessage            string `json:"last_message"`
	SendTime               int64  `json:"last_message_time"`
	FromType               string `json:"from_type"`
	LastMessageIsAutoReply bool   `json:"last_message_is_auto_reply"`
	BotStopAt              int64  `json:"bot_stop_at"`
	OpSource               string `json:"op_source"`
	OpSourceSendTime       int64  `json:"op_source_send_time"`
}

func (t Thread) TableName() string {
	return "threads"
}

func MapperToThread(thread *response.Thread) Thread {
	return Thread{
		SQLModel:            *core.NewUpsertWithoutIdSQLModel(),
		ChannelCode:         thread.ChannelCode,
		PlatformThreadId:    thread.ThreadId,
		PlatformCustomerId:  thread.CustomerId,
		CustomerName:        thread.CustomerName,
		EncodedCustomerName: thread.CustomerName,
		CustomerAvatarUrl:   thread.CustomerAvatarUrl,
		UnreadCount:         thread.UnreadCount,
		PlatformCode:        thread.Platform,
	}
}
