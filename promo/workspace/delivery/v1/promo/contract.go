package promo

import (
	"pensiel.com/domain/src/usecase/promo"
	"pensiel.com/material/src/contract"
)

type (
	GetPromoListResponseData struct {
		SecondaryId string `json:"secondaryId"`
		Name        string `json:"name"`
		Discount    uint64 `json:"discount"`
		Code        string `json:"code"`
	}

	GetPromoListResponse struct {
		contract.ResponseMeta
		Data []GetPromoListResponseData `json:"data"`
	}
)

func (res *GetPromoListResponseData) MapFromModel(model promo.PromoListReturn) {
	res.SecondaryId = model.SecondaryId
	res.Name = model.Name
	res.Discount = model.Discount
	res.Code = model.Code
}

type (
	GetPromoDetailResponseData struct {
		SecondaryId string `json:"secondaryId"`
		Name        string `json:"name"`
		Discount    uint64 `json:"discount"`
		Code        string `json:"code"`
	}

	GetPromoDetailResponse struct {
		contract.ResponseMeta
		Data GetPromoDetailResponseData `json:"data"`
	}
)

func (res *GetPromoDetailResponseData) MapFromModel(model promo.PromoDetailReturn) {
	res.SecondaryId = model.SecondaryId
	res.Name = model.Name
	res.Discount = model.Discount
	res.Code = model.Code
}
