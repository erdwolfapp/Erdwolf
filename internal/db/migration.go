package db

const MigrationLogTableName = "ed_db_migrate_log"
const VerbMigrationReverted = "down"
const VerbMigrationApplied = "up"

type DBChangeSchema struct {
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

func findSchemaById(list []DBChangeSchema, id string) (*DBChangeSchema, bool) {
	for _, object := range list {
		if object.Info.Identifier == id {
			return &object, true
		}
	}

	return nil, false
}

type DBMigrateLog struct {
	gorm.Model

	Identifier		string	`gorm:"size:64"`
	Verb			string 	`gorm:"size:4"`
}

func (DBMigrateLog) TableName() string {
	return MigrationLogTableName
}

func findLatestSchemaLogEntry(list []DBMigrateLog, id string) (*DBMigrateLog, bool) {
	// Find latest DBMigrateLog by ID in a typical log array.
	// Useful in case a migration was reverted by system admin.
	for index := len(list) - 1; index >= 0; index-- {
		object := &list[index]

		if object.Identifier == id {
			return object, true
		}
	}

	return nil, false
}