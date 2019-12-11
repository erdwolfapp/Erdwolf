package frontend

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (t *FrontendHandlers) ServeSignIn(c echo.Context) error {
	if !t.app.IsAnyAuthProviderAvailable() {
		return errors.New("no auth providers are configured")
	}

	return c.Render(http.StatusOK, "sign_in.html", map[string]interface{}{
		"isDevelopment": t.app.IsDevelopment(),
		"configuredAuthProviders": t.app.GetAuthSubDomainList(),
	})
}
