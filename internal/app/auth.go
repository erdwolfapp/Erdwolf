package app

import (
	"errors"
	"fmt"
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

func (a *Application) RegisterAuthDomainFactory(factory auth.DomainFactory) {
	a.authDomainFactories[factory.DomainId()] = factory
}

func (a *Application) NewAuthDomain(name string, config auth.DomainConfig) error {
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

func (a *Application) GetAuthSubDomainList() *map[string]auth.Domain {
	return &a.authDomains
}