package chat_mapper

import (
	"jetshop/services/chat_service/internal/modules/chat_model"
	"jetshop/shared/common/enums"
	"jetshop/shared/integration/hermes/response"
	"jetshop/shared/sctx/core"
)

func MapperToThread(thread *response.Thread) (*chat_model.Thread, error) {
	platformCode, err := enums.ParsePlatformCode(thread.Platform)
	if err != nil {
		return nil, err
	}

	return &chat_model.Thread{
		SQLModel:           *core.NewUpsertWithoutIdSQLModel(),
		ChannelCode:        thread.ChannelCode,
		PlatformThreadId:   thread.ThreadId,
		PlatformCustomerId: thread.CustomerId,
		CustomerName:       thread.CustomerName,
		//EncodedCustomerName: thread.CustomerName,
		CustomerAvatarUrl: thread.CustomerAvatarUrl,
		UnreadCount:       thread.UnreadCount,
		PlatformCode:      platformCode,
	}, nil
}
