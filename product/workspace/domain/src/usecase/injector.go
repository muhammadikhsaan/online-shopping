package usecase

import (
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase/product"
	"pensiel.com/material/src/client"
)

type Service struct {
	Product product.Usecase
}

func NewService(r *data.Repository, c *client.Client) *Service {
	return &Service{
		Product: product.NewService(c, &product.Repository{
			Product: r.Product,
		}),
	}
}
