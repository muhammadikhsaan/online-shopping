package product

import (
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/pensiel"
)

type Handler interface {
	Router(r pensiel.Router)
}

type handler struct {
	uc *usecase.Service
}

func NewHandler(uc *usecase.Service) Handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) Router(r pensiel.Router) {
	r.Get("/product", h.GETPRODUCTLIST)
	r.Get("/product/{secondaryId}", h.GETPRODUCTDETAIL)
}
