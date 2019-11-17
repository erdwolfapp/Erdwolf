package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"html/template"
)

type ErdwolfServer struct {
	echo *echo.Echo
	config ApplicationConfig
}

func NewAppServer(config ApplicationConfig) *ErdwolfServer {
	e := echo.New()
	server := &ErdwolfServer {
		echo: e,
		config: config,
	}
	server.echo.HideBanner = true

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/templates/*.gohtml", server.config.ResourcesPath))),
	}
	e.Renderer = renderer
	server.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} #${id}, remote:${remote_ip}, status:${status}, ${method}:${uri}\n",
	}))
	server.echo.Use(middleware.Recover())

	server.setupRoutes()
	return server
}

func (s *ErdwolfServer) Start() {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Http.Port)); err != nil {
		s.echo.Logger.Fatal(errors.Wrap(err, "failed to start the application server"))
	}
}