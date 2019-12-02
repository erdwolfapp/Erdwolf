package app

func (a *Application) IsDevelopment() bool {
	return a.appConfig.Application.Environment == "development"
}

func (a *Application) GetResourcesPath() string {
	return a.appConfig.Resources.Path
}