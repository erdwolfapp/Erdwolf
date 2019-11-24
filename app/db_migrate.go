package app

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
)

const DB_MIGRATE_LOG_TABLE_NAME = "ed_db_migrate_log"

type DBMigrateLog struct {
	gorm.Model

	Identifier		string	`gorm:"unique;primary_key"`
	Verb			string 	`gorm:"size:4"`
}

func (DBMigrateLog) TableName() string {
	return DB_MIGRATE_LOG_TABLE_NAME
}

type DBMigrationSchema struct {
	Info struct {
		Identifier	string
		Description	string
	}
	Sql  struct {
		Up			string
		Down		string
	}
}

func (a *Application) MigrateDatabase() error {
	fmt.Println("==> Running database migration...")
	if !a.orm.HasTable(DB_MIGRATE_LOG_TABLE_NAME) {
		fmt.Println("===> Migration log is not initialized yet. Initializing...")
		fmt.Println(a.orm.AutoMigrate(DBMigrateLog{}))
	}

	// Pull the migration log
	var migrationLog []DBMigrateLog
	a.orm.Find(&migrationLog)

	// Discover available migration schemas
	availableMigrations := map[string]DBMigrationSchema{}
	err := filepath.Walk("./migrations", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		var schema DBMigrationSchema
		if _, err := toml.DecodeFile(path, &schema); err != nil {
			return err
		}
		availableMigrations[schema.Info.Identifier] = schema
		return nil
	})
	if err != nil {
		return err
	}

	for _, logEntry := range migrationLog {
		if _, exists := availableMigrations[logEntry.Identifier]; !exists {
			return errors.New(fmt.Sprintf("Missing schema info for %s", logEntry.Identifier))
		}
	}

	return errors.New("fixme: TODO: Implement MigrateDatabase()")
}