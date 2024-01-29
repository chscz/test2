package http

import (
	"ab180/internal/domain"
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type ShortLinkRepo interface {
	CreateShortLink(ctx context.Context, shortID domain.ShortLink) error
	GetURLByShortLink(ctx context.Context, shortLink string) (domain.ShortLink, error)
}
type ShortLinkHandler struct {
	repo        ShortLinkRepo
	wAPI        api.WriteAPI
	measurement string
}

func NewShortLinkHandler(repo ShortLinkRepo, wAPI api.WriteAPI, measurement string) *ShortLinkHandler {
	return &ShortLinkHandler{
		repo:        repo,
		wAPI:        wAPI,
		measurement: measurement,
	}
}
