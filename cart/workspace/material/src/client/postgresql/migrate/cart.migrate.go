package migrate

import (
	"gorm.io/gorm"
	"pensiel.com/domain/src/data/cart"
	"pensiel.com/material/src/client/postgresql/executor"
)

type cartmigrator struct {
	*gorm.DB
}

func Cart(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&cartmigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *cartmigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (mg *cartmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(cart.EntityModel{}); !ok {
		if err := m.CreateTable(cart.EntityModel{}); err != nil {
			return err
		}

		if err := mg.Seeder(m); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*cartmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(cart.EntityModel{}); ok {
		if err := m.DropTable(cart.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*cartmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*cartmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*cartmigrator) Index(m *executor.Migrator) error {
	return nil
}

// Seeder implements executor.Actions.
func (*cartmigrator) Seeder(m *executor.Migrator) error {
	return nil
}
