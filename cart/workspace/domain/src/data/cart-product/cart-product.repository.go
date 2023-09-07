package cartproduct

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
	Delete(ctx context.Context, model *EntityModel) *pensiel.Error
	FindByIndex(ctx context.Context, model *EntityModel) *pensiel.Error
	Insert(ctx context.Context, model *EntityModel) *pensiel.Error
	InsertOrUpdateByCartIdAndProductId(ctx context.Context, model *EntityModel) *pensiel.Error
	GetAllByIndex(ctx context.Context, model *EntityModel) ([]EntityModel, *pensiel.Error)
}

type repository struct {
	dbx helper.DatabaseExtractionFunc
}

func NewRepository(dbx helper.DatabaseExtractionFunc) Repository {
	return &repository{
		dbx: dbx,
	}
}

func (r *repository) InsertOrUpdateByCartIdAndProductId(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	dbi = dbi.
		Model(model).
		Where("cart_id = ?", model.CartId).
		Where("product_id = ?", model.ProductId).
		Update("quantity", model.Quantity)

	if err := dbi.Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed update product cart into server",
			Origin:     err,
		}
	}

	if dbi.RowsAffected == 0 {
		if err := r.Insert(ctx, model); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) Insert(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Create(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed insert new product cart into server",
			Origin:     err,
		}
	}

	return nil
}

func (r *repository) FindByIndex(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Where(model).First(model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "cart product not found",
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

func (r *repository) GetAllByIndex(ctx context.Context, model *EntityModel) ([]EntityModel, *pensiel.Error) {
	dbi := r.dbx(ctx)

	data := []EntityModel{}

	if err := dbi.Where(model).Find(&data).Error; err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed get cart into server",
			Origin:     err,
		}
	}

	return data, nil
}

func (r *repository) Update(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Save(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed update cart into server",
			Origin:     err,
		}
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := r.dbx(ctx)

	if err := dbi.Delete(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed delete cart into server",
			Origin:     err,
		}
	}

	return nil
}
