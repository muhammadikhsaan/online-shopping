package v1

import (
	"pensiel.com/delivery/v1/promo"
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/pensiel"
)

type Delivery interface {
	Router(r pensiel.Router)
}

type delivery struct {
	uc *usecase.Service
}

func NewDelivery(uc *usecase.Service) Delivery {
	return &delivery{
		uc: uc,
	}
}

func (c *delivery) Router(r pensiel.Router) {
	// USERS
	r.Group(promo.NewHandler(c.uc).Router)
}
