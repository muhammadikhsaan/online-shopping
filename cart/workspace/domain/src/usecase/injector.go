package usecase

import (
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase/cart"
	"pensiel.com/material/src/client"
)

type Service struct {
	Cart cart.Usecase
}

func NewService(r *data.Repository, c *client.Client) *Service {
	return &Service{
		Cart: cart.NewService(c, &cart.Repository{
			Product:     r.Product,
			Cart:        r.Cart,
			CartProduct: r.CartProduct,
			Promo:       r.Promo,
		}),
	}
}
