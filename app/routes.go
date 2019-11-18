package app

import (
	"fmt"
	"github.com/erdwolfapp/Erdwolf/frontend"
)

func (s *ErdwolfServer) setupRoutes() {
	// Home page
	s.echo.GET("/", frontend.ServeIndex)

	// Static resources
	s.echo.Static("/static/cooked", fmt.Sprintf("%s/cooked", s.config.ResourcesPath))
}