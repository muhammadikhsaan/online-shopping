package migrate

import (
	"gorm.io/gorm"
	cartproduct "pensiel.com/domain/src/data/cart-product"
	"pensiel.com/material/src/client/postgresql/executor"
)

type cartproductmigrator struct {
	*gorm.DB
}

func CartProduct(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&cartproductmigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *cartproductmigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (mg *cartproductmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(cartproduct.EntityModel{}); !ok {
		if err := m.CreateTable(cartproduct.EntityModel{}); err != nil {
			return err
		}

		if err := mg.Seeder(m); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*cartproductmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(cartproduct.EntityModel{}); ok {
		if err := m.DropTable(cartproduct.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*cartproductmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*cartproductmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*cartproductmigrator) Index(m *executor.Migrator) error {
	return nil
}

// Seeder implements executor.Actions.
func (*cartproductmigrator) Seeder(m *executor.Migrator) error {
	return nil
}
