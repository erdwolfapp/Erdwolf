package app

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
)

func dbMigrateFindSchemas() ([]DBChangeSchema, error) {
	var results []DBChangeSchema
	err := filepath.Walk("./migrations", func (path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		schema := DBChangeSchema{}
		if _, err := toml.DecodeFile(path, &schema); err != nil {
			return err
		}

		results = append(results, schema)
		return nil
	})

	return results, err
}

func (a *Application) MigrateDatabase() error {
	fmt.Println("==> Running database migration...")
	if !a.orm.HasTable(DB_MIGRATE_LOG_TABLE_NAME) {
		fmt.Println("===> Migration log is not initialized yet. Initializing...")
		a.orm.AutoMigrate(DBMigrateLog{})
		if a.orm.Error != nil {
			return a.orm.Error
		}
	}

	// Pull the migration log
	fmt.Println("===> Retrieving logs...")
	var changeLog []DBMigrateLog
	a.orm.Find(&changeLog)

	// Discover available migration schemas
	fmt.Println("===> Discovering local migration data...")
	localSchemas, err := dbMigrateFindSchemas()
	if err != nil {
		return err
	}

	// Check if we have info about applied migrations
	fmt.Println("===> Validating past data...")
	for _, logEntry := range changeLog {
		schema, found := findSchemaById(localSchemas, logEntry.Identifier)
		if !found {
			return errors.New(fmt.Sprintf("\"%s\" has been applied but its schema is missing", logEntry.Identifier))
		}

		if logEntry.Verb != VerbMigrationApplied && logEntry.Verb != VerbMigrationReverted {
			return errors.New(fmt.Sprintf("unsafe to continue due to corruption: no state info for change \"%s\"", logEntry.Identifier))
		}

		if logEntry.Verb == VerbMigrationReverted && !schema.Info.SafeApplied {
			fmt.Printf("WARN: Potentially unsafe change was applied before: %s - %s\n", schema.Info.Identifier, schema.Info.Description)
		}
	}

	fmt.Println("===> Checking whether structures can be updated...")
	for _, schema := range localSchemas {
		if !schema.Info.SafeApplied {
			fmt.Printf("   Skipping a potentially unsafe change: %s (code: \"%s\")\n", schema.Info.Identifier, schema.Info.Identifier)
			continue
		}

		log, found := findLatestSchemaLogEntry(changeLog, schema.Info.Identifier)
		if found && log.Verb == VerbMigrationApplied {
			continue
		}

		fmt.Printf("+  Applying change: %s\n", schema.Info.Description)
		a.orm.Exec(schema.Sql.Up)
		if a.orm.Error != nil {
			return a.orm.Error
		}

		// Mark the change as merged
		newLog := DBMigrateLog {
			Identifier: schema.Info.Identifier,
			Verb: VerbMigrationApplied,
		}
		a.orm.Create(&newLog)
		if a.orm.Error != nil {
			return a.orm.Error
		}
	}

	return nil
}