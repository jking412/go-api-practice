package migrate

import (
	"go-api-practice/pkg/console"
	"go-api-practice/pkg/database"
	"go-api-practice/pkg/file"
	"gorm.io/gorm"
	"io/ioutil"
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

func (migrator *Migrator) createMigrationsTable() {
	migration := &Migration{}

	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up() {
	migrateFiles := migrator.readAllMigrationFiles()

	batch := migrator.getBatch()

	// 获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	runed := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrateFiles {

		// 对比文件名称，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

func (migrator *Migrator) getBatch() int {
	batch := 1
	lastMigration := Migration{}

	migrator.DB.Order("id DESC").First(&lastMigration)

	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {

	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrationFiles []MigrationFile

	for _, f := range files {
		fileName := file.FileNameWithoutExtension(f.Name())

		mfile := getMigrationFile(fileName)

		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrationFiles, mfile)
		}
	}

	return migrationFiles
}
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	if mfile.Up != nil {
		console.Warning("migrating " + mfile.FileName)
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		console.Success("migrated " + mfile.FileName)
	}
	err := database.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error
	console.ExitIf(err)
}
