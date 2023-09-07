package promo

import (
	"net/http"

	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func (h *handler) GETPROMOLIST(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	responses := []GetPromoListResponseData{}

	promos, err := h.uc.Promo.PromoList(ctx)
	if err != nil {
		return err
	}

	for _, promo := range promos {
		response := GetPromoListResponseData{}
		response.MapFromModel(promo)
		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, GetPromoListResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "success to get promo data",
		},
		Data: responses,
	})
}

func (h *handler) GETPROMODETAIL(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	promoCode := c.Param("promoCode")

	responses := GetPromoDetailResponseData{}
	promo, err := h.uc.Promo.PromoDetail(ctx, promoCode)

	if err != nil {
		return err
	}

	responses.MapFromModel(*promo)

	return c.JSON(http.StatusOK, GetPromoDetailResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "success to get promo data",
		},
		Data: responses,
	})
}
