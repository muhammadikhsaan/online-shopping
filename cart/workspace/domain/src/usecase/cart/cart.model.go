package cart

import (
	"sync"

	"pensiel.com/domain/src/data/cart"
	"pensiel.com/domain/src/data/product"
)

type (
	AddProductIntoCartParam struct {
		Invoice   string
		ProductId string
		Quantity  uint32
	}
)

type (
	UpdateQuantityProductCartParam struct {
		Invoice   string
		ProductId string
		Quantity  uint32
	}
)

type (
	CartPromoParam struct {
		Invoice   string
		PromoCode string
	}
)

type (
	RemoveProductFormCartParam struct {
		Invoice   string
		ProductId string
	}
)

type (
	GetCartDetailParam struct {
		Invoice string
	}

	GetProductCartDataProduct struct {
		SecondaryId     string `json:"secondaryId"`
		SKU             string `json:"sku"`
		Name            string `json:"string"`
		DisplayQuantity string `json:"displayQuatity"`
		Quatity         uint32 `json:"quatity"`
		Price           uint64 `json:"price"`
		TotalPrice      uint64 `json:"totalPrice"`
		Unit            string `json:"unit"`
	}

	GetCartDetailData struct {
		SecondaryId   string                      `json:"secondaryId"`
		Invoice       string                      `json:"invoice"`
		Price         uint64                      `json:"price"`
		DiscountPrice uint64                      `json:"discountPrice"`
		FinalPrice    uint64                      `json:"finalPrice"`
		PromoCode     string                      `json:"promoCode"`
		Product       []GetProductCartDataProduct `json:"product"`
	}

	GetCartDetailSync struct {
		mu   sync.Mutex
		cart GetCartDetailData
	}

	GetCartDetailReturn = GetCartDetailData
)

func (d *GetCartDetailData) MapFromEntityCart(entity *cart.EntityModel) {
	d.SecondaryId = entity.SecondaryId
	d.Invoice = entity.Invoice
	d.PromoCode = entity.PromoCode
}

func (d *GetProductCartDataProduct) MapFromEntityProduct(entity *product.ResponseData) {
	d.SecondaryId = entity.SecondaryId
	d.DisplayQuantity = entity.DisplayQuantity
	d.Name = entity.Name
	d.SKU = entity.SKU
	d.Quatity = entity.Quatity
	d.Price = entity.Price
	d.Unit = entity.Unit
}
