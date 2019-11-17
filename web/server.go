package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"net/http"
)

type ErdwolfServer struct {
	echo *echo.Echo
	config ApplicationConfig
}

func NewAppServer(config ApplicationConfig) *ErdwolfServer {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} #${id}, remote:${remote_ip}, status:${status}, ${method}:${uri}\n",
	}))
	e.Use(middleware.Recover())
	e.GET("/", hello)

	return &ErdwolfServer {
		echo: e,
		config: config,
	}
}

func (s *ErdwolfServer) Start() {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Http.Port)); err != nil {
		s.echo.Logger.Fatal(errors.Wrap(err, "failed to start the application server"))
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}