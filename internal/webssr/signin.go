package webssr

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Module) ServeSignIn(c echo.Context) error {
	if !m.app.IsAnyAuthProviderAvailable() {
		return errors.New("no auth providers are configured")
	}

	return c.Render(http.StatusOK, "sign_in.html", map[string]interface{}{
		"isDevelopment": m.app.IsDevelopment(),
		"configuredAuthProviders": m.app.GetAuthSubDomainList(),
	})
}
