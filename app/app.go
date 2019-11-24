package app

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Application struct {
	appConfig			ErdwolfConfig

	authDomains			map[string]AuthDomain
	authDomainFactories map[string]AuthDomainFactory
	orm					*gorm.DB
	http				*echo.Echo
}

func NewInstance(appConfig ErdwolfConfig) Application {
	return Application {
		appConfig: appConfig,

		authDomains: map[string]AuthDomain{},
		authDomainFactories: map[string]AuthDomainFactory{},
		orm: nil,
		http: nil,
	}
}

func (a *Application) CleanUp() error {
	if err := a.orm.Close(); err != nil {
		return err
	}

	return nil
}