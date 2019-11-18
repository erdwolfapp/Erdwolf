package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/erdwolfapp/Erdwolf/app"
)

func main() {
	appConfig := app.ApplicationConfig{
		Environment: "development",
		ResourcesPath: "resources",
		Http: app.HttpConfig {
			Port: 8080,
		},
	}

	dbConfig := app.DatabaseConfig{}

	if _, err := toml.DecodeFile("config/application.toml", &appConfig); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := toml.DecodeFile("config/database.toml", &dbConfig); err != nil {
		fmt.Println(err)
		return
	}

	erdwolf := app.NewInstance(appConfig, dbConfig)
	if err := erdwolf.InitHttpServer(); err != nil {
		fmt.Println(err)
		return
	}
	erdwolf.StartListening()
}