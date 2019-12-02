package main

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/app"
	"github.com/erdwolfapp/Erdwolf/app/auth/nullauth"
)

func configureAuthImplementations(erdwolf *app.Application) {
	fmt.Println("Registering auth domain factories")
	erdwolf.RegisterAuthDomainFactory(nullauth.NewFactory())
	//erdwolf.RegisterAuthDomainFactory(oauth.NewFactory())
}

func configureAuth(erdwolf *app.Application) {
	for subdomainId, subdomainConfig := range appConfig.AuthDomains {
		fmt.Printf("Configuring auth provider: %s\n", subdomainId)
		if err := erdwolf.NewAuthDomain(subdomainId, subdomainConfig); err != nil {
			handleError(err)
		}
	}
}