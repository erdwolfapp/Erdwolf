package main

import "github.com/erdwolfapp/Erdwolf/internal/configs"

var appConfig = configs.ErdwolfConfig{
	Application: configs.ApplicationPrivateConfig {
		Environment: "development",
		HttpPort: 8080,
	},

	Resources: configs.ResourcesPrivateConfig {
		Path: "resources",
	},

	AuthDomains: configs.AuthDomainDefs{},
	Secrets: configs.SecretsPrivateConfig{},
}