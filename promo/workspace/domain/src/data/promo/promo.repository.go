package promo

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
	"pensiel.com/material/src/helper"
	"pensiel.com/material/src/pensiel"
)

type Repository interface {
	FindAll(ctx context.Context) ([]EntityModel, *pensiel.Error)
	FindByIndex(ctx context.Context, model *EntityModel) *pensiel.Error
}

type repository struct {
	dbx helper.DatabaseExtractionFunc
}

func NewRepository(dbx helper.DatabaseExtractionFunc) Repository {
	return &repository{
		dbx: dbx,
	}
}

func (r *repository) FindAll(ctx context.Context) ([]EntityModel, *pensiel.Error) {
	dbi := r.dbx(ctx)

	model := []EntityModel{}

	err := dbi.
		Find(&model).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model, nil
		}

		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed get promo list from sistem",
			Origin:     err,
		}
	}

	return model, nil
}

func (r *repository) FindByIndex(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	err := dbi.
		Where(model).
		First(&model).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "promo not found",
				Origin:     err,
			}
		}

		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed get promo from sistem",
			Origin:     err,
		}
	}

	return nil
}
