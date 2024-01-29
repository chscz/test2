package http

import "github.com/labstack/echo"

func SetRoute(e *echo.Echo, h *ShortLinkHandler) *echo.Echo {
	e.GET("/", h.Index)
	e.POST("/short-links", h.CreateShortLink)

	e.GET("/short-links", h.GetShortLink)
	e.GET("/r/:short-link", h.Redirect)

	return e
}
