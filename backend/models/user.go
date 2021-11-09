package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`

	LikedNews     []*News `gorm:"many2many:users_news"`
	RecommendNews []*News `gorm:"many2many:recommend_news"`
}
