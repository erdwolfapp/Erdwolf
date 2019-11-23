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
	Create(AuthDomainConfig) AuthDomain
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

func (a *Application) NewAuthDomain(name string, config AuthDomainConfig) error {
	// Set sub-domain ID to "name" if none is specified.
	if config.SubDomainId == "" {
		config.SubDomainId = name
	}

	implementation, exists := a.authDomainFactories[config.CoreImplementation]
	if !exists {
		return errors.New(fmt.Sprintf("no underlaying implementation found for %s/%s", config.CoreImplementation, config.SubDomainId))
	}

	instance := implementation.Create(config)
	a.authDomains[instance.DomainId()] = instance
	return nil
}