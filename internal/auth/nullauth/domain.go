package nullauth

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

type NullAuthDomain struct {
	config auth.DomainConfig
}

func (d *NullAuthDomain) Name() string {
	if d.config.FriendlyName == "" {
		return "no validation"
	}

	return d.config.FriendlyName
}

func (d *NullAuthDomain) DomainId() string {
	return fmt.Sprintf("null/%s", d.config.SubDomainId)
}