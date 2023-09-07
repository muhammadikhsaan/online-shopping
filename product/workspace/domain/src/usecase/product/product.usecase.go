package product

import (
	"context"

	"pensiel.com/domain/src/data/product"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

type Usecase interface {
	ProductList(ctx context.Context) ([]ProductListReturn, *pensiel.Error)
	ProductDetail(ctx context.Context, secondaryId string) (*ProductDetailReturn, *pensiel.Error)
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	Product product.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) ProductList(ctx context.Context) ([]ProductListReturn, *pensiel.Error) {
	dbx := uc.Dbi.Cnx(ctx)

	products, err := uc.Product.FindAll(dbx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc *usecase) ProductDetail(ctx context.Context, secondaryId string) (*ProductDetailReturn, *pensiel.Error) {
	dbx := uc.Dbi.Cnx(ctx)

	entity := &product.EntityModel{
		MetaEntity: contract.MetaEntity{
			ShowableEntity: contract.ShowableEntity{
				SecondaryId: secondaryId,
			},
		},
	}

	if err := uc.Product.FindByIndex(dbx, entity); err != nil {
		return nil, err
	}

	return entity, nil
}
