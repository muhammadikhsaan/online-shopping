package promo

import "pensiel.com/material/src/contract"

type (
	ResponseData struct {
		SecondaryId string `json:"secondaryId"`
		Name        string `json:"name"`
		Discount    uint64 `json:"discount"`
		Code        string `json:"code"`
	}

	Response struct {
		contract.ResponseMeta
		Data ResponseData `json:"data"`
	}
)
