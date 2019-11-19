package app

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/frontend"
)

func (a *Application) setupRoutes() {
	// Public: Home page / Log in
	a.http.GET("/", frontend.ServeIndex)
	a.http.GET("/auth/sign-in", frontend.ServeSignIn)

	// Restricted: Dashboard

	// Static resources
	a.http.Static("/static/cooked", fmt.Sprintf("%s/cooked", a.appConfig.ResourcesPath))
}