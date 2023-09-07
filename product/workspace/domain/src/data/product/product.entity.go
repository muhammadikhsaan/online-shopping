package product

import (
	"pensiel.com/material/src/contract"
)

type Entity struct {
	SKU     string `json:"sku"`
	Name    string `json:"string"`
	Quatity uint32 `json:"quatity"`
	Price   uint32 `json:"price"`
	Unit    string `json:"unit"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "products"
}
