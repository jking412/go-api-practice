package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(migrator gorm.Migrator, db *sql.DB)

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
func getMigrationFile(name string) MigrationFile {
	for _, migrationFile := range migrationFiles {
		if migrationFile.FileName == name {
			return migrationFile
		}
	}
	return MigrationFile{}
}

func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
