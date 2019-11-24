package app

import (
	"github.com/labstack/echo/v4"
)

type Application struct {
	appConfig			ErdwolfConfig

	authDomains			map[string]AuthDomain
	authDomainFactories map[string]AuthDomainFactory
	http				*echo.Echo
}

func NewInstance(appConfig ErdwolfConfig) Application {
	return Application {
		appConfig: appConfig,

		authDomains: map[string]AuthDomain{},
		authDomainFactories: map[string]AuthDomainFactory{},
		http: nil,
	}
}