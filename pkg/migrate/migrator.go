package migrate

import (
	"go-api-practice/pkg/database"
	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primarykey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	migrator.createMigrationsTable()

	return migrator
}

func (m *Migrator) createMigrationsTable() {
	migration := &Migration{}

	if !m.Migrator.HasTable(&migration) {
		m.Migrator.CreateTable(&migration)
	}
}
