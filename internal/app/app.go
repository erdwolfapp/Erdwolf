package app

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/erdwolfapp/Erdwolf/internal/configs"
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

type Application struct {
	appConfig			configs.ErdwolfConfig

	authDomains			map[string]auth.Domain
	authDomainFactories map[string]auth.DomainFactory
	orm					*gorm.DB
	http				*echo.Echo
}

func NewInstance(appConfig ErdwolfConfig) Application {
	return Application {
		appConfig: appConfig,

		authDomains: map[string]auth.Domain{},
		authDomainFactories: map[string]auth.DomainFactory{},
		orm: nil,
		http: nil,
	}
}

func (a *Application) CleanUp() error {
	fmt.Println("==> Running shutdown tasks")
	if err := a.closeDBConnection(); err != nil {
		fmt.Printf("====> Error: %v\n", err)
	}
	return nil
}