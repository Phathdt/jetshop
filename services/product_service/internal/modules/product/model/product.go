package model

import (
	"jetshop/service-context/core"
)

type Product struct {
	core.SQLModel
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (Product) TableName() string {
	return "products"
}
