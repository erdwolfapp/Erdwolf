package app

type AuthProviders 	= map[string]interface{}
type AuthDomainDefs = map[string]AuthDomainConfig
type SecretsConfig 	= map[string]string

type ApplicationConfig struct {
	Environment 	string 			`toml:"application.environment"`
	ResourcesPath	string			`toml:"application.resourcesPath"`

	AuthDomains		AuthDomainDefs	`toml:"application.auth"`
	Http 			HttpConfig 		`toml:"application.http"`
	Secrets			SecretsConfig	`toml:"application.secrets"`
}

type HttpConfig struct {
	Port 			int `toml:"port"`
}

func (s *ApplicationConfig) IsDevelopment() bool {
	return s.Environment == "development"
}

func (a *Application) GetResourcesPath() string {
	return a.appConfig.ResourcesPath
}