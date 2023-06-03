package channel_model

import (
	"jetshop/shared/sctx/core"
)

type Channel struct {
	core.SQLModel
	Name   string `json:"name"`
	Code   string `json:"code"`
	Active bool   `json:"active"`
}

func (Channel) TableName() string {
	return "onpoint.channels"
}
