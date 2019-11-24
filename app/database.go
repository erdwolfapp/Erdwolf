package app

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (a *Application) OpenDBConnection() error {
	// TODO: Rename a.appConfig.Database.Driver to a.appConfig.Database.Dialect
	// TODO: Missing null-check for a.appConfig.Database.Driver

	if a.appConfig.Database.Driver != "sqlite3" {
		return errors.New("currently database dialects other than SQLite are not supported")
	}

	// TODO: Missing null-check for a.appConfig.Database.SQLite.Path
	db, err := gorm.Open(a.appConfig.Database.Driver, a.appConfig.Database.SQLite.Path)
	if err != nil {
		return err
	}

	a.orm = db
	return nil
}