package auth

type AuthDomain struct {
	DomainId string
	Name string
}

type AuthConfig struct {
	Use 				string `toml:"use"`
	SubDomainId 		string `toml:"domainId"`
	FriendlyName 		string `toml:"friendlyName"`
	ServiceDomain 		string `toml:"serviceDomain"`
	SecretsNamespace 	string `toml:"secrets"`
}