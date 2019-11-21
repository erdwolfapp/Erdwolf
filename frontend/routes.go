package frontend

import (
	"fmt"
)

func (t *FrontendHandlers) InstallRoutes() {
	http := t.app.GetHttp()

	// Public: Home page / Log in
	http.GET("/", t.ServeIndex)
	http.GET("/auth/sign-in", t.ServeSignIn)

	// Restricted: Dashboard

	// Static resources
	http.Static("/static/cooked", fmt.Sprintf("%s/cooked", t.app.GetResourcesPath()))
}