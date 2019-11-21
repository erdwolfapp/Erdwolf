package app

import (
	"errors"
	"fmt"
)

type AuthDomain interface {
	Name() string
	DomainId() string
}

type AuthDomainFactory interface {
	DomainId() string
	Create(AuthConfig) *AuthDomain
}

type ExtraAuthData = map[string]string

type AuthDomainConfig struct {
	CoreImplementation	string 			`toml:"coreImplementation"`
	SubDomainId 		string 			`toml:"domainId"`
	FriendlyName 		string 			`toml:"friendlyName"`
	SecretsNamespace 	string 			`toml:"secrets"`
	ExtraData			ExtraAuthData	`toml:"extra"`
}

func (a *Application) RegisterAuthDomainFactory(factory AuthDomainFactory) {
	a.authDomainFactories[factory.DomainId()] = factory
}

func (a *Application) NewAuthDomain(config AuthDomainConfig) error {
	implementation, exists := a.authDomainFactories[config.CoreImplementation]

	if !exists {
		return errors.New(fmt.Sprintf("no underlaying implementation found for %s/%s", config.CoreImplementation, config.SubDomainId))
	}

	// TODO: Do something with the auth domain instance.
	fmt.Println("fixme:todo: Auth domain created but __not registered__.")
	implementation.Create(config)
	return nil
}