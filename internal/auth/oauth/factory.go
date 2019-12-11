package oauth

import "github.com/erdwolfapp/Erdwolf/internal/auth"

type OAuthDomainFactory struct {}

func (f *OAuthDomainFactory) DomainId() string {
	return "oauth"
}

func (f *OAuthDomainFactory) Create(config auth.DomainConfig) auth.Domain {
	return &OAuthDomain {
		config: config,
	}
}

func NewFactory() auth.DomainFactory {
	return &OAuthDomainFactory{}
}