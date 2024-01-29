package http

import (
	"github.com/labstack/echo"
	"net/http"
)

func (sh *ShortLinkHandler) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "index",
	})
}
