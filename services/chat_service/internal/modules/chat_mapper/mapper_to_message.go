package chat_mapper

import (
	"jetshop/services/chat_service/internal/modules/chat_enums"
	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/common/enums"
	"jetshop/shared/integration/hermes/response"
	"jetshop/shared/sctx/core"
)

func MapperToMessage(message *response.Message) (*chat_model.Message, error) {
	messageType, err := chat_enums.ParseMessageType(message.MessageType)
	if err != nil {
		messageType = chat_enums.MessageTypeUnknown
	}

	fromType, err := chat_enums.ParseFromType(message.FromType)
	if err != nil {
		fromType = chat_enums.FromTypeSystem
	}

	var status chat_enums.MessageStatus

	switch message.Status {
	case "normal":
		status = chat_enums.MessageStatusReceived
	case "deleted":
		status = chat_enums.MessageStatusDeleted
	default:
		status = chat_enums.MessageStatusSent
	}

	platformCode, err := enums.ParsePlatformCode(message.Platform)
	if err != nil {
		return nil, err
	}

	return &chat_model.Message{
		SQLModel:          *core.NewUpsertWithoutIdSQLModel(),
		ChannelCode:       message.ChannelCode,
		PlatformCode:      platformCode,
		PlatformThreadId:  message.ThreadId,
		PlatformMessageId: message.MessageId,
		Content:           message.Content,
		MessageType:       messageType,
		FromType:          fromType,
		SendTime:          message.SendTime,
		Status:            status,
		IsAutoReply:       message.AutoReply,
	}, nil
}
