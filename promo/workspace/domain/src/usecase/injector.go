package usecase

import (
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase/promo"
	"pensiel.com/material/src/client"
)

type Service struct {
	Promo promo.Usecase
}

func NewService(r *data.Repository, c *client.Client) *Service {
	return &Service{
		Promo: promo.NewService(c, &promo.Repository{
			Promo: r.Promo,
		}),
	}
}
