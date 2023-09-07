package product

import (
	"net/http"

	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func (h *handler) GETPRODUCTLIST(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	responses := []GetProductListResponseData{}

	products, err := h.uc.Product.ProductList(ctx)
	if err != nil {
		return err
	}

	for _, product := range products {
		response := GetProductListResponseData{}
		response.MapFromModel(product)
		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, GetProductListResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "success to get product data",
		},
		Data: responses,
	})
}

func (h *handler) GETPRODUCTDETAIL(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	secondaryId := c.Param("secondaryId")

	responses := GetProductDetailResponseData{}
	product, err := h.uc.Product.ProductDetail(ctx, secondaryId)

	if err != nil {
		return err
	}

	responses.MapFromModel(*product)

	return c.JSON(http.StatusOK, GetProductDetailResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "success to get product data",
		},
		Data: responses,
	})
}
