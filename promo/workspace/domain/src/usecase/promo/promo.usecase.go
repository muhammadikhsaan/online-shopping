package promo

import (
	"context"

	"pensiel.com/domain/src/data/promo"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/pensiel"
)

type Usecase interface {
	PromoList(ctx context.Context) ([]PromoListReturn, *pensiel.Error)
	PromoDetail(ctx context.Context, promoCode string) (*PromoDetailReturn, *pensiel.Error)
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	Promo promo.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) PromoList(ctx context.Context) ([]PromoListReturn, *pensiel.Error) {
	dbx := uc.Dbi.Cnx(ctx)

	promos, err := uc.Promo.FindAll(dbx)

	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (uc *usecase) PromoDetail(ctx context.Context, promoCode string) (*PromoDetailReturn, *pensiel.Error) {
	dbx := uc.Dbi.Cnx(ctx)

	entity := &promo.EntityModel{
		Entity: promo.Entity{
			Code: promoCode,
		},
	}

	if err := uc.Promo.FindByIndex(dbx, entity); err != nil {
		return nil, err
	}

	return entity, nil
}
