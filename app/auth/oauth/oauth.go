package app

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/app/auth"
	"golang.org/x/oauth2"
)

type OAuthConfig struct {
	Provider string
	Domain string
	AuthorizationUri string
}

func NewOAuth2Domain(config OAuthConfig) auth.AuthDomain {
	return auth.AuthDomain {
		DomainId: fmt.Sprintf("oauth/%s", config.Provider),
	}
}
