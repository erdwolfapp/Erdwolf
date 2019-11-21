package app

func (a *Application) IsAnyAuthProviderAvailable() bool {
	return len(a.authDomains) > 0
}