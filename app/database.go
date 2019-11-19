package app

import (
	"errors"
	"fmt"
)

type DatabaseConfig struct {
	EnableAutoMigrations	bool	`toml:"database.enable-auto-migrations"`
	Driver		 			string 	`toml:"database.driver"`
	// SQLite
	Path					string	`toml:"database.sqlite.path"`
}

func (a *Application) MigrateDatabase() error {
	fmt.Println("==> Running database migration...")
	return errors.New("fixme: TODO: Implement MigrateDatabase()")
}