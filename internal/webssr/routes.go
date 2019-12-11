package webssr

import (
	"fmt"
)

import "github.com/erdwolfapp/Erdwolf/internal/app"

type Module struct {
	app *app.Application
}

func New(app *app.Application) Module {
	return Module {
		app: app,
	}
}

func (m *Module) InstallRoutes() {
	http := m.app.GetHttp()

	// Public: Home page / Log in
	http.GET("/", m.ServeIndex)
	http.GET("/auth/sign-in", m.ServeSignIn)

	// Restricted: Dashboard

	// Static resources
	http.Static("/static/cooked", fmt.Sprintf("%s/cooked", m.app.GetResourcesPath()))
}