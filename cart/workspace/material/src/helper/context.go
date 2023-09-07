package helper

import (
	"context"

	"gorm.io/gorm"
	"pensiel.com/material/src/client/postgresql"
)

type DatabaseExtractionFunc func(ctx context.Context) *gorm.DB

func DatabaseExtraction(ctx context.Context) *gorm.DB {
	return ctx.(*postgresql.Connection).Conn
}
