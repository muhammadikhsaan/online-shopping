package cart

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
	"pensiel.com/material/src/helper"
	"pensiel.com/material/src/pensiel"
)

type Repository interface {
	Update(ctx context.Context, model *EntityModel) *pensiel.Error
	Insert(ctx context.Context, model *EntityModel) *pensiel.Error
	FindByInvoice(ctx context.Context, model *EntityModel) *pensiel.Error
}

type repository struct {
	dbx helper.DatabaseExtractionFunc
}

func NewRepository(dbx helper.DatabaseExtractionFunc) Repository {
	return &repository{
		dbx: dbx,
	}
}

func (r *repository) Insert(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Create(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed insert new cart into server",
			Origin:     err,
		}
	}

	return nil
}

func (r *repository) FindByInvoice(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	dbi = dbi.Where(model).First(model)

	if err := dbi.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "cart not found",
				Origin:     err,
			}
		}

		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed get cart into server",
			Origin:     err,
		}
	}

	return nil
}

func (r *repository) Update(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Save(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed get cart into server",
			Origin:     err,
		}
	}

	return nil
}
