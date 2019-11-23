package nullauth

import "github.com/erdwolfapp/Erdwolf/app"

type NullAuthDriver struct {}

func (f *NullAuthDriver) DomainId() string {
	return "null"
}

func (f *NullAuthDriver) Create(config app.AuthDomainConfig) app.AuthDomain {
	return &NullAuthDomain {
		config: config,
	}
}

func NewFactory() app.AuthDomainFactory {
	return &NullAuthDriver{}
}