package main

import "github.com/erdwolfapp/Erdwolf/app"

var appConfig = app.ErdwolfConfig{
	Application: app.ApplicationPrivateConfig {
		Environment: "development",
		HttpPort: 8080,
	},

	Resources: app.ResourcesPrivateConfig {
		Path: "resources",
	},

	AuthDomains: app.AuthDomainDefs{},
	Secrets: app.SecretsConfig{},
}