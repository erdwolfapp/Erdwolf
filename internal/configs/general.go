package configs

import (
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

type ErdwolfConfig struct {
	Application 	ApplicationPrivateConfig
	Database		DatabasePrivateConfig
	Resources		ResourcesPrivateConfig

	AuthDomains		auth.DomainDefs	`toml:"authentication"`
	Secrets			SecretsPrivateConfig
}