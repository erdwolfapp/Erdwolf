package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := fmt.Sprintf("error-%d.gohtml", code)
	if err2 := c.Render(code, errorPage, nil); err2 != nil {
		c.Logger().Error(err2)
	}
	c.Logger().Error(err)
}