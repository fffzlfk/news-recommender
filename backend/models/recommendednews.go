package models

type RecommendedNews struct {
	UserID        uint  `gorm:"primaryKey"`
	NewsID        uint  `gorm:"primaryKey"`
	RecommendedAt int64 `json:"recommended_at"`
}
