package configs

type DatabasePrivateConfig struct {
	EnableAutoMigrations	bool	`toml:"autoMigrate"`
	Driver		 			string

	SQLite					SQLitePrivateConfig
}

type SQLitePrivateConfig struct {
	Path		string
}