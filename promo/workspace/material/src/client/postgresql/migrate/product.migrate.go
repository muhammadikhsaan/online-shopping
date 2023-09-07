package migrate

import (
	"gorm.io/gorm"
	"pensiel.com/domain/src/data/promo"
	"pensiel.com/material/src/client/postgresql/executor"
	"pensiel.com/material/src/contract"
)

type promomigrator struct {
	*gorm.DB
}

func Promo(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&promomigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *promomigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (mg *promomigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(promo.EntityModel{}); !ok {
		if err := m.CreateTable(promo.EntityModel{}); err != nil {
			return err
		}

		if err := mg.Seeder(m); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*promomigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(promo.EntityModel{}); ok {
		if err := m.DropTable(promo.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*promomigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*promomigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*promomigrator) Index(m *executor.Migrator) error {
	return nil
}

// Seeder implements executor.Actions.
func (*promomigrator) Seeder(m *executor.Migrator) error {
	m.Insert(&[]promo.EntityModel{
		{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: "2301n2iyu3c4xn78912y",
				},
			},
			Entity: promo.Entity{
				Name:     "Special 17 Agustus",
				Discount: 25,
				Code:     "SPC1708",
			},
		},
	})
	return nil
}
