package app

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/frontend"
)

func (a *Application) setupRoutes() {
	// Home page
	a.http.GET("/", frontend.ServeIndex)

	// Static resources
	a.http.Static("/static/cooked", fmt.Sprintf("%s/cooked", a.appConfig.ResourcesPath))
}