package data

import (
	"pensiel.com/domain/src/data/promo"
	"pensiel.com/material/src/helper"
)

type Repository struct {
	Promo promo.Repository
}

func NewRepository() *Repository {
	dbx := helper.DatabaseExtraction

	promo := promo.NewRepository(dbx)

	return &Repository{
		Promo: promo,
	}
}
