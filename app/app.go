package app

import "github.com/labstack/echo/v4"

type Application struct {
	appConfig		ApplicationConfig
	databaseConfig	DatabaseConfig

	http			*echo.Echo
}

func NewInstance(appConfig ApplicationConfig, databaseConfig DatabaseConfig) Application {
	return Application {
		appConfig: appConfig,
		databaseConfig: databaseConfig,

		http: nil,
	}
}