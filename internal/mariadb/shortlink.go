package mariadb

import (
	"ab180/internal/domain"
	"context"
)

func (r *ShortLinkRepo) CreateShortLink(ctx context.Context, shortLink domain.ShortLink) error {
	return r.DB.WithContext(ctx).Create(shortLink).Error
}
func (r *ShortLinkRepo) GetURLByShortLink(ctx context.Context, shortLinkID string) (domain.ShortLink, error) {
	var shortLink domain.ShortLink
	if err := r.DB.WithContext(ctx).Where("id = ?", shortLinkID).Take(&shortLink).Error; err != nil {
		return domain.ShortLink{}, err
	}
	return shortLink, nil
}
