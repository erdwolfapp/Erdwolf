package configs

type ErdwolfConfig struct {
	Application 	ApplicationPrivateConfig
	Database		DatabasePrivateConfig
	Resources		ResourcesPrivateConfig

	AuthDomains		AuthDomainDefs			`toml:"authentication"`
	Secrets			SecretsPrivateConfig
}