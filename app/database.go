package app

type DatabaseConfig struct {
	EnableAutoMigrations	bool	`toml:"database.enable-auto-migratons"`
	Driver		 			string 	`toml:"database.driver"`
	// SQLite
	Path					string	`toml:"database.sqlite.path"`
}

