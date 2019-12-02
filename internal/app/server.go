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
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/templates/*.gohtml", a.GetResourcesPath()))),
	}

	a.http = echo.New()
	a.http.HideBanner = true
	a.http.HTTPErrorHandler = httpErrorHandler
	a.http.Renderer = renderer

	// Middlewares
	a.http.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} remote:${remote_ip}, status:${status}, ${method}:${uri}\n",
	}))
	a.http.Use(middleware.Recover())

	// Session middleware
	sessionsSecret := a.GetSecret("COOKIE_STORE_SECRET", "ChangeMe!!!")
	a.http.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionsSecret))))

	return nil
}

func (a *Application) GetHttp() *echo.Echo {
	return a.http
}

func (a *Application) StartListening() {
	if err := a.http.Start(fmt.Sprintf(":%d", a.appConfig.Application.HttpPort)); err != nil {
		a.http.Logger.Fatal(errors.Wrap(err, "Failed to start the application server"))
	}
}