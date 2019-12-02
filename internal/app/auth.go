package app

import (
	"errors"
	"fmt"
)

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

func (a *Application) IsAnyAuthProviderAvailable() bool {
	return len(a.authDomains) > 0
}

func (a *Application) GetAuthSubDomainList() *map[string]AuthDomain {
	return &a.authDomains
}