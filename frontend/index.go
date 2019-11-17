package frontend

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServeIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.gohtml", map[string]interface{}{
	})
}
