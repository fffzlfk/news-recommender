package models

type LikedNews struct {
	UserID  uint  `gorm:"primaryKey"`
	NewsID  uint  `gorm:"primaryKey"`
	LikedAt int64 `json:"liked_at"`
}
