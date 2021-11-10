package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`

	LikedNews     []News `gorm:"many2many:liked_news"`
	RecommendNews []News `gorm:"many2many:recommend_news"`
}
