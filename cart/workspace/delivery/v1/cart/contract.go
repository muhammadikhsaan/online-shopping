package cart

import (
	"pensiel.com/domain/src/usecase/cart"
	"pensiel.com/material/src/contract"
)

type (
	GetProductCartResponse struct {
		contract.ResponseMeta
		Data cart.GetCartDetailData `json:"data"`
	}
)

type (
	AddProductIntoCartRequest struct {
		ProductId string `json:"productId" validate:"required"`
		Quantity  uint32 `json:"quantity" validate:"required,number"`
	}

	AddProductIntoCartResponse struct {
		contract.ResponseMeta
	}
)

type (
	UpdateQuantityProductCartRequest struct {
		Quantity uint32 `json:"quantity" validate:"required,number"`
	}

	UpdateQuantityProductCartResponse struct {
		contract.ResponseMeta
	}
)

type (
	ApplyPromoRequest struct {
		PromoCode string `json:"promoCode" validate:"required"`
	}

	ApplyPromoResponse struct {
		contract.ResponseMeta
	}
)
