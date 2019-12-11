package oauth

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

type OAuthDomain struct {
	config auth.DomainConfig
}

func (d *OAuthDomain) Name() string {
	if d.config.FriendlyName == "" {
		return "an OAuth2 provider"
	}

	return d.config.FriendlyName
}

func (d *OAuthDomain) DomainId() string {
	return fmt.Sprintf("null/%s", d.config.SubDomainId)
}