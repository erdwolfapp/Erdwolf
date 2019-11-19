package frontend

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServeSignIn(c echo.Context) error {
	return c.Render(http.StatusOK, "sign_in.gohtml", map[string]interface{}{
	})
}
