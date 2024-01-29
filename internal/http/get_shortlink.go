package http

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func (sh *ShortLinkHandler) GetShortLink(c echo.Context) error {
	inputShortLinkID := c.FormValue("inputShortLinkID")
	if strings.HasPrefix(inputShortLinkID, baseURL) {
		inputShortLinkID = strings.ReplaceAll(inputShortLinkID, baseURL, "")
	}
	if inputShortLinkID == "" {
		return c.String(http.StatusBadRequest, "invalid link id")
	}

	shortLink, err := sh.repo.GetURLByShortLink(context.Background(), inputShortLinkID)
	if err != nil {
		return c.String(http.StatusBadRequest, "not exist link id")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": map[string]string{
			"shortId":   shortLink.ID,
			"url":       shortLink.URL,
			"createdAt": shortLink.CreatedAt.Time.Format(RespTimeFormat),
		},
	})
}
