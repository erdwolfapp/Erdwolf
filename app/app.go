package app

import (
	"github.com/labstack/echo/v4"
)

type Application struct {
	appConfig			ApplicationConfig
	databaseConfig		DatabaseConfig

	authDomains			[]*AuthDomain
	authDomainFactories map[string]AuthDomainFactory
	http				*echo.Echo
}

func NewInstance(appConfig ApplicationConfig, databaseConfig DatabaseConfig) Application {
	return Application {
		appConfig: appConfig,
		databaseConfig: databaseConfig,

		authDomains: []*AuthDomain{},
		authDomainFactories: map[string]AuthDomainFactory{},
		http: nil,
	}
}