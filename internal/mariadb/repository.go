package mariadb

import "gorm.io/gorm"

type ShortLinkRepo struct {
	DB *gorm.DB
}
