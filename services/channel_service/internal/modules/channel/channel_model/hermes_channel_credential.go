package channel_model

import (
	"jetshop/shared/proto/out/proto"
	"jetshop/shared/sctx/core"
)

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

func (c *HermesChannelCredential) ToProtoc() *jetshop_proto.HermesChannelCredential {
	return &jetshop_proto.HermesChannelCredential{
		ChannelCode:  c.ChannelCode,
		PlatformCode: c.PlatformCode,
		IsEnabled:    c.IsEnabled,
		SellerId:     c.SellerId,
	}
}
