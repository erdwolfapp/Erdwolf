package app

type AuthProviders 	= map[string]interface{}
type AuthDomainDefs = map[string]AuthDomainConfig
type SecretsConfig 	= map[string]string

type ErdwolfConfig struct {
	Application 	ApplicationPrivateConfig
	Database		DatabasePrivateConfig
	Resources		ResourcesPrivateConfig

	AuthDomains		AuthDomainDefs	`toml:"authentication"`
	Secrets			SecretsConfig
}

type ApplicationPrivateConfig struct {
	Environment string
	HttpPort	int
}

type ResourcesPrivateConfig struct {
	Path		string
}

type DatabasePrivateConfig struct {
	EnableAutoMigrations	bool	`toml:"autoMigrate"`
	Driver		 			string

	SQLite					SQLitePrivateConfig
}

type SQLitePrivateConfig struct {
	Path		string
}

func (a *Application) IsDevelopment() bool {
	return a.appConfig.Application.Environment == "development"
}

func (a *Application) GetResourcesPath() string {
	return a.appConfig.Resources.Path
}