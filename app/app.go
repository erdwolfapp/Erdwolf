package app

import (
	"github.com/labstack/echo/v4"
)

type Application struct {
	appConfig			ErdwolfConfig
	databaseConfig		DatabaseConfig

	authDomains			map[string]AuthDomain
	authDomainFactories map[string]AuthDomainFactory
	http				*echo.Echo
}

func NewInstance(appConfig ErdwolfConfig, databaseConfig DatabaseConfig) Application {
	return Application {
		appConfig: appConfig,
		databaseConfig: databaseConfig,

		authDomains: map[string]AuthDomain{},
		authDomainFactories: map[string]AuthDomainFactory{},
		http: nil,
	}
}