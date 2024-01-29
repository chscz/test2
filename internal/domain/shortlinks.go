package domain

import (
	"database/sql"
)

type ShortLink struct {
	ID        string       `gorm:"column:id;primaryKey"`
	CreatedAt sql.NullTime `gorm:"column:created_at"`
	URL       string       `gorm:"column:url"`
}

func (ShortLink) TableName() string {
	return "short_link"
}
