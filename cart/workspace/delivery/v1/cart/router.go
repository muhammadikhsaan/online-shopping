package cart

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
	r.Get("/cart/{invoice}", h.GETPRODUCTCART)
	r.Post("/cart/{invoice}", h.ADDPRODUCTINTOCART)
	r.Post("/cart/{invoice}/promo", h.APPLYPROMO)
	r.Patch("/cart/{invoice}/{productId}", h.UPDATEQUATITYPRODUCTCART)
	r.Delete("/cart/{invoice}/{productId}", h.REMOVEPRODUCTCART)
}
