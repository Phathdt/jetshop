package channel_model

import "jetshop/service-context/core"

type HermesChannelCredential struct {
	core.SQLModel
	ChannelCode  string `json:"channel_code"`
	PlatformCode string `json:"platform_code"`
	ExpiredAt    int    `json:"expired_at"`
	IsEnabled    bool   `json:"is_enabled"`
	SellerId     string `json:"seller_id"`
}

func (HermesChannelCredential) TableName() string {
	return "hermes.hermes_channel_credentials"
}
