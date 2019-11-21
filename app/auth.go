package app

func (a *Application) IsAnyAuthProviderAvailable() bool {
	return len(a.AppConfig.Auth) > 0
}