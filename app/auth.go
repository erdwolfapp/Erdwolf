package app

func (a *Application) IsAnyAuthProviderAvailable() bool {
	return len(a.authDomains) > 0
}

func (a *Application) GetAuthSubDomainList() *map[string]AuthDomain {
	return &a.authDomains
}