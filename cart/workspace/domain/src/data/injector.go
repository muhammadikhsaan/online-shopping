package data

import (
	"pensiel.com/domain/src/data/cart"
	cartproduct "pensiel.com/domain/src/data/cart-product"
	"pensiel.com/domain/src/data/product"
	"pensiel.com/domain/src/data/promo"
	"pensiel.com/material/src/helper"
)

type Repository struct {
	Product     product.Repository
	Promo       promo.Repository
	Cart        cart.Repository
	CartProduct cartproduct.Repository
}

func NewRepository() *Repository {
	dbx := helper.DatabaseExtraction

	product := product.NewRepository()
	promo := promo.NewRepository()
	cart := cart.NewRepository(dbx)
	cartproduct := cartproduct.NewRepository(dbx)

	return &Repository{
		Product:     product,
		Promo:       promo,
		Cart:        cart,
		CartProduct: cartproduct,
	}
}
