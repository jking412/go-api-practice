package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(migrator *gorm.Migrator, db *sql.DB)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
