package auth

type Domain interface {
	Name() string
	DomainId() string
}

type DomainFactory interface {
	DomainId() string
	Create(DomainConfig) Domain
}

type ExtraData = map[string]string
type DomainConfig struct {
	CoreImplementation	string 			`toml:"coreImplementation"`
	SubDomainId 		string 			`toml:"domainId"`
	FriendlyName 		string 			`toml:"friendlyName"`
	SecretsNamespace 	string 			`toml:"secrets"`
	ExtraData			ExtraData	`toml:"extra"`
}