package db

import "github.com/jinzhu/gorm"

const MigrationLogTableName = "ed_db_migrate_log"
const VerbMigrationReverted = "down"
const VerbMigrationApplied = "up"

type SchemaChange struct {
	Info struct {
		Identifier	string
		Description	string
		SafeApplied	bool
	}
	Sql  struct {
		Up			string
		Down		string
	}
}

func FindSchemaById(list []SchemaChange, id string) (*SchemaChange, bool) {
	for _, object := range list {
		if object.Info.Identifier == id {
			return &object, true
		}
	}

	return nil, false
}

type SchemaChangeLog struct {
	gorm.Model

	Identifier		string	`gorm:"size:64"`
	Verb			string 	`gorm:"size:4"`
}

func (SchemaChangeLog) TableName() string {
	return MigrationLogTableName
}

func FindLatestSchemaLogEntry(list []SchemaChangeLog, id string) (*SchemaChangeLog, bool) {
	// Find latest SchemaChangeLog by ID in a typical log array.
	// Useful in case a migration was reverted by system admin.
	for index := len(list) - 1; index >= 0; index-- {
		object := &list[index]

		if object.Identifier == id {
			return object, true
		}
	}

	return nil, false
}