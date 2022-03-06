package models

import "gorm.io/gorm"

type RecommendedNews struct {
	gorm.Model

	UserID        uint  `gorm:"primaryKey"`
	NewsID        uint  `gorm:"primaryKey"`
	RecommendedAt int64 `json:"recommended_at"`
}
