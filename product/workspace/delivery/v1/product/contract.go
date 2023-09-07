package product

import (
	"fmt"

	"pensiel.com/domain/src/usecase/product"
	"pensiel.com/material/src/contract"
)

type (
	GetProductListResponseData struct {
		SecondaryId     string `json:"secondaryId"`
		SKU             string `json:"sku"`
		Name            string `json:"string"`
		DisplayQuantity string `json:"displayQuatity"`
		Quatity         uint32 `json:"quatity"`
		Price           uint32 `json:"price"`
		Unit            string `json:"unit"`
	}

	GetProductListResponse struct {
		contract.ResponseMeta
		Data []GetProductListResponseData `json:"data"`
	}
)

func (res *GetProductListResponseData) MapFromModel(model product.ProductListReturn) {
	res.SecondaryId = model.SecondaryId
	res.Name = model.Name
	res.Price = model.Price
	res.Quatity = model.Quatity
	res.DisplayQuantity = fmt.Sprintf("%d%s", model.Quatity, model.Unit)
	res.SKU = model.SKU
	res.Unit = model.Unit
}

type (
	GetProductDetailResponseData struct {
		SecondaryId     string `json:"secondaryId"`
		SKU             string `json:"sku"`
		Name            string `json:"string"`
		DisplayQuantity string `json:"displayQuatity"`
		Quatity         uint32 `json:"quatity"`
		Price           uint32 `json:"price"`
		Unit            string `json:"unit"`
	}

	GetProductDetailResponse struct {
		contract.ResponseMeta
		Data GetProductDetailResponseData `json:"data"`
	}
)

func (res *GetProductDetailResponseData) MapFromModel(model product.ProductDetailReturn) {
	res.SecondaryId = model.SecondaryId
	res.Name = model.Name
	res.Price = model.Price
	res.Quatity = model.Quatity
	res.DisplayQuantity = fmt.Sprintf("%d%s", model.Quatity, model.Unit)
	res.SKU = model.SKU
	res.Unit = model.Unit
}
