package app

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"html/template"
)

func (a *Application) InitHttpServer() error {
	if a.http != nil {
		return errors.New("Application server is already initialized.")
	}

	a.http = echo.New()
	a.http.HideBanner = true

	// Template Engine
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/templates/*.gohtml", a.appConfig.ResourcesPath))),
	}
	a.http.Renderer = renderer

	// Middlewares
	a.http.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} #${id}, remote:${remote_ip}, status:${status}, ${method}:${uri}\n",
	}))
	a.http.Use(middleware.Recover())

	// Session middleware
	sessionsSecret := []byte(a.appConfig.Secrets.Sessions)
	a.http.Use(session.Middleware(sessions.NewCookieStore(sessionsSecret)))

	a.setupRoutes()
	return nil
}

func (a *Application) StartListening() {
	if err := a.http.Start(fmt.Sprintf(":%d", a.appConfig.Http.Port)); err != nil {
		a.http.Logger.Fatal(errors.Wrap(err, "Failed to start the application server"))
	}
}