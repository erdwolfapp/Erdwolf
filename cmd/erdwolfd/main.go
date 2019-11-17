package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/erdwolfapp/Erdwolf/app"
)

func main() {
	config := app.ApplicationConfig{
		Environment: "development",
		ResourcesPath: "resources",
		Http: app.HttpConfig {
			Port: 8080,
		},
	}
	if _, err := toml.DecodeFile("config/application.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	server := app.NewAppServer(config)
	server.Start()
}