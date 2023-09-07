package cart

import (
	"pensiel.com/material/src/contract"
)

type Entity struct {
	Invoice   string `json:"invoice"`
	PromoCode string `json:"promoCode"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "cart"
}
