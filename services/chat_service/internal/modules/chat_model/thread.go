package chat_model

import (
	"database/sql/driver"
	"encoding/json"

	"jetshop/services/chat_service/internal/modules/chat_enums"
	"jetshop/shared/common/enums"
	"jetshop/shared/sctx/core"
)

type Thread struct {
	core.SQLModel
	ChannelCode            string              `json:"channel_code"`
	PlatformThreadId       string              `json:"platform_thread_id"`
	PlatformCustomerId     string              `json:"platform_customer_id"`
	CustomerName           string              `json:"customer_name"`
	CustomerAvatarUrl      string              `json:"customer_avatar_url"`
	UnreadCount            int                 `json:"unread_count"`
	PlatformCode           enums.PlatformCode  `json:"platform_code"`
	LastMessage            *MessageContent     `json:"last_message"`
	SendTime               int64               `json:"last_message_time"`
	FromType               chat_enums.FromType `json:"from_type"`
	LastMessageIsAutoReply bool                `json:"last_message_is_auto_reply"`
	//EncodedCustomerName    string              `json:"encoded_customer_name"`
	//BotStopAt              int64               `json:"bot_stop_at"`
	//OpSource               string              `json:"op_source"`
	//OpSourceSendTime       int64               `json:"op_source_send_time"`
}

func (t Thread) TableName() string {
	return "threads"
}

type MessageContent struct {
	Content     string                 `json:"content,omitempty"`
	SendTime    int64                  `json:"sendTime,omitempty"`
	MessageType chat_enums.MessageType `json:"message_type,omitempty"`
}

func (m *MessageContent) Value() (driver.Value, error) {
	val, err := json.Marshal(m)

	return string(val), err
}

func (m *MessageContent) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), &m)
}
