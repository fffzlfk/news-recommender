package models

type News struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url" gorm:"UNIQUE"`
	UrlToImage  string `json:"url_to_image"`
	Category    string `json:"category"`
	CreatedAt   int64  `json:"created_at"`

	BeLikedBy []*User `gorm:"many2many:users_news"`
}

var Categorys = []string{"business", "entertainment", "general",
	"health", "science", "sports", "technology"}
