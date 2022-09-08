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

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback() {

	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	// 回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 回退迁移，按照倒序执行迁移的 down 方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {

	// 标记是否真的有执行了迁移回退的操作
	runed := false

	for _, _migration := range migrations {

		// 友好提示
		console.Warning("rollback " + _migration.Migration)

		// 执行迁移文件的 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true

		// 回退成功了就删除掉这条记录
		migrator.DB.Delete(&_migration)

		// 打印运行状态
		console.Success("finish " + mfile.FileName)
	}
	return runed
}

// Reset 回滚所有迁移
func (migrator *Migrator) Reset() {

	migrations := []Migration{}

	// 按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

// Refresh 回滚所有迁移，并运行所有迁移
func (migrator *Migrator) Refresh() {

	// 回滚所有迁移
	migrator.Reset()

	// 再次执行所有迁移
	migrator.Up()
}

// Fresh Drop 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {

	// 获取数据库名称，用以提示
	dbname := database.CurrentDatabase()

	// 删除所有表
	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Success("clearup database " + dbname)

	// 重新创建 migrates 表
	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	migrator.Up()
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
