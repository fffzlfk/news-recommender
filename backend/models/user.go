package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`

	LikedNews       []News `gorm:"many2many:liked_news"`
	ClickedNews     []News `gorm:"many2many:clicked_news"`
	RecommendedNews []News `gorm:"many2many:recommended_news"`
}
