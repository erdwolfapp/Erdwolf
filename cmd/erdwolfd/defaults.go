package main

import "github.com/erdwolfapp/Erdwolf/app"

var appConfig = app.ApplicationConfig{
	Environment: "development",
	ResourcesPath: "resources",
	Http: app.HttpConfig {
		Port: 8080,
	},
}

var dbConfig = app.DatabaseConfig{}