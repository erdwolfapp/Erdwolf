package configs

import (
	"github.com/erdwolfapp/Erdwolf/internal/auth"
)

type AuthProviders 	= map[string]interface{}
type AuthDomainDefs = map[string]auth.DomainConfig