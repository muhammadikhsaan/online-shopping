package promo

import (
	"pensiel.com/material/src/contract"
)

type Entity struct {
	Name     string `json:"name"`
	Discount uint64 `json:"discount"`
	Code     string `json:"code"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "promo"
}
