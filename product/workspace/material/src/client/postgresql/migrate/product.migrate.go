package migrate

import (
	"gorm.io/gorm"
	"pensiel.com/domain/src/data/product"
	"pensiel.com/material/src/client/postgresql/executor"
	"pensiel.com/material/src/contract"
)

type productsmigrator struct {
	*gorm.DB
}

func Product(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&productsmigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *productsmigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (mg *productsmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(product.EntityModel{}); !ok {
		if err := m.CreateTable(product.EntityModel{}); err != nil {
			return err
		}

		if err := mg.Seeder(m); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*productsmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(product.EntityModel{}); ok {
		if err := m.DropTable(product.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*productsmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*productsmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*productsmigrator) Index(m *executor.Migrator) error {
	return nil
}

// Seeder implements executor.Actions.
func (*productsmigrator) Seeder(m *executor.Migrator) error {
	m.Insert(&[]product.EntityModel{
		{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: "aynGK823985300PsnLH",
				},
			},
			Entity: product.Entity{
				Name:    "mie instan",
				SKU:     "MI9087",
				Quatity: 100,
				Price:   3200,
				Unit:    "PCS",
			},
		},
		{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: "aynGK823985300PsnLZ",
				},
			},
			Entity: product.Entity{
				Name:    "air mineral",
				SKU:     "AM9086",
				Quatity: 76,
				Price:   3000,
				Unit:    "PCS",
			},
		},
		{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: "aynGK823985300PsnLp",
				},
			},
			Entity: product.Entity{
				Name:    "susu murni",
				SKU:     "SM7652",
				Quatity: 55,
				Price:   6300,
				Unit:    "PCS",
			},
		},
		{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: "aynGK823985305PsnLp",
				},
			},
			Entity: product.Entity{
				Name:    "roti tawar",
				SKU:     "SM7882",
				Quatity: 0,
				Price:   15000,
				Unit:    "PCS",
			},
		},
	})
	return nil
}
