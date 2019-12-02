package oauth

import "github.com/erdwolfapp/Erdwolf/app"

type OAuthDomainFactory struct {}

func (f *OAuthDomainFactory) DomainId() string {
	return "oauth"
}

func (f *OAuthDomainFactory) Create(config app.AuthDomainConfig) app.AuthDomain {
	return &OAuthDomain {}
}

func NewFactory() app.AuthDomainFactory {
	return &OAuthDomainFactory{}
}