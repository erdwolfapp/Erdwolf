package nullauth

import "github.com/erdwolfapp/Erdwolf/internal/auth"

type NullAuthDriver struct {}

func (f *NullAuthDriver) DomainId() string {
	return "null"
}

func (f *NullAuthDriver) Create(config auth.DomainConfig) auth.Domain {
	return &NullAuthDomain {
		config: config,
	}
}

func NewFactory() auth.DomainFactory {
	return &NullAuthDriver{}
}