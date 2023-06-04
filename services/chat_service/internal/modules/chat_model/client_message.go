package chat_model

import (
	"jetshop/services/chat_service/internal/modules/chat_enums"
	"jetshop/shared/common/enums"
)

type ClientMessage struct {
	Id                int                            `json:"id"`
	SendTime          int64                          `json:"send_time"`
	SentByUserId      int                            `json:"sent_by_user_id"`
	Content           string                         `json:"content"`
	PlatformMessageId string                         `json:"platform_message_id"`
	MessageType       chat_enums.ClientMessageType   `json:"message_type"`
	FromType          chat_enums.FromType            `json:"from_type"`
	Status            chat_enums.ClientMessageStatus `json:"status"`
	PlatformId        string                         `json:"platform_id"`
	Platform          enums.PlatformCode             `json:"platform"`
}
