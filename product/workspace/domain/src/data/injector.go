package data

import (
	"pensiel.com/domain/src/data/product"
	"pensiel.com/material/src/helper"
)

type Repository struct {
	Product product.Repository
}

func NewRepository() *Repository {
	dbx := helper.DatabaseExtraction

	product := product.NewRepository(dbx)

	return &Repository{
		Product: product,
	}
}
