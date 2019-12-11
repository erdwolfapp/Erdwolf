package webssr

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Module) ServeIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
	})
}
