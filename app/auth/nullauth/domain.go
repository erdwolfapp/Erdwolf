package nullauth

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/app"
)

type NullAuthDomain struct {
	config app.AuthDomainConfig
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