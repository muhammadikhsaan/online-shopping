package product

import "pensiel.com/material/src/contract"

type (
	ResponseData struct {
		SecondaryId     string `json:"secondaryId"`
		SKU             string `json:"sku"`
		Name            string `json:"string"`
		DisplayQuantity string `json:"displayQuatity"`
		Quatity         uint32 `json:"quatity"`
		Price           uint64 `json:"price"`
		Unit            string `json:"unit"`
	}

	Response struct {
		contract.ResponseMeta
		Data ResponseData `json:"data"`
	}
)
