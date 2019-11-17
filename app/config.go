package app

type HttpConfig struct {
	Port int `toml:"port"`
}

type ApplicationConfig struct {
	Environment 	string 		`toml:"application.environment"`
	ResourcesPath	string		`toml:"application.resourcesPath"`
	Http 			HttpConfig 	`toml:"application.http"`
}

func (s *ApplicationConfig) IsDevelopment() bool {
	return s.Environment == "development"
}