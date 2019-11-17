package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/erdwolfapp/Erdwolf/web"
)

func main() {
	config := web.ApplicationConfig{
		Environment: "development",
		Http: web.HttpConfig {
			Port: 8080,
		},
	}
	if _, err := toml.DecodeFile("config/application.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	server := web.NewAppServer(config)
	server.Start()
}