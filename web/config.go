package web

type HttpConfig struct {
	Port int `toml:"port"`
}

type ErdwolfConfig struct {
	Environment string 		`toml:"application.environment"`
	Http 		HttpConfig 	`toml:"application.http"`
}

func (s *ErdwolfConfig) IsDevelopment() bool {
	return s.Environment == "development"
}