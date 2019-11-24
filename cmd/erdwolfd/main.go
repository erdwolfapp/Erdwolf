package main

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/app"
	"github.com/erdwolfapp/Erdwolf/frontend"
)

func handleError(err error) {
	fmt.Printf("Caught an error: %+v\n", err)
}

func main() {
	if err := loadConfig(&appConfig, CONFIG_APPLICATION); err != nil {
		handleError(err)
		return
	}

	erdwolf := app.NewInstance(appConfig)

	if dbConfig.EnableAutoMigrations {
		if err := erdwolf.MigrateDatabase(); err != nil {
			handleError(err)
			return
		}
	}

	configureAuthImplementations(&erdwolf)
	configureAuth(&erdwolf)

	if err := erdwolf.InitHttpServer(); err != nil {
		handleError(err)
		return
	}
	front := frontend.New(&erdwolf)
	front.InstallRoutes()
	erdwolf.StartListening()
}