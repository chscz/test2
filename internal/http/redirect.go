package http

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func (sh *ShortLinkHandler) Redirect(c echo.Context) error {
	ctx := context.Background()
	shortLinkID := c.Param("short-link")
	shortLink, err := sh.repo.GetURLByShortLink(ctx, shortLinkID)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid link id")
	}

	p := influxdb2.NewPoint("data",
		map[string]string{"id": shortLink.ID},
		map[string]interface{}{"url": shortLink.URL},
		time.Now().UTC())
	sh.wAPI.WritePoint(p)
	// Flush writes
	sh.wAPI.Flush()

	return c.Redirect(http.StatusFound, shortLink.URL)
}
