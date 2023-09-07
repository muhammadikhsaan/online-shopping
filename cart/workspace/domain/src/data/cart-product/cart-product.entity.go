package cartproduct

import "pensiel.com/material/src/contract"

type Entity struct {
	CartId    string `json:"cartId"`
	ProductId string `json:"productId"`
	Quantity  uint32 `json:"quantity"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "cart-product"
}
