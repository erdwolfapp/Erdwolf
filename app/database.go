package app

import (
	"errors"
	"fmt"
)

func (a *Application) MigrateDatabase() error {
	fmt.Println("==> Running database migration...")
	return errors.New("fixme: TODO: Implement MigrateDatabase()")
}