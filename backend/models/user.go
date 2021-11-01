package models

import (
	"errors"
	"news-api/utils/weightedrandom"

	"github.com/spf13/viper"
)

type User struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email" gorm:"unique"`
	Password         []byte `json:"-"`
	VisBusiness      int    `json:"vis_business" gorm:"default:1"`
	VisEntertainment int    `json:"vis_entertainment" gorm:"default:1"`
	VisGeneral       int    `json:"vis_general" gorm:"default:1"`
	VisHealth        int    `json:"vis_health" gorm:"default:1"`
	VisScience       int    `json:"vis_scientce" gorm:"default:1"`
	VisSports        int    `json:"vis_sports" gorm:"default:1"`
	VisTechnology    int    `json:"vis_technology" gorm:"default:1"`

	LikedNews []News `gorm:"many2many:users_news"`
}

func (user *User) GetANewsCategory() (string, error) {
	wrc, err := weightedrandom.NewChooser(map[string]int{
		"business":      user.VisBusiness,
		"entertainment": user.VisEntertainment,
		"general":       user.VisGeneral,
		"health":        user.VisEntertainment,
		"science":       user.VisScience,
		"sports":        user.VisSports,
		"technology":    user.VisTechnology,
	})
	if err != nil {
		return "", err
	}

	result := wrc.Pick()
	return result, nil
}

var increasement int = viper.GetInt("number.increasement")

func (user *User) LikeNews(news *News, do bool) error {
	var inc int
	if do {
		inc = int(increasement)
	} else {
		inc = -int(increasement)
	}
	switch news.Category {
	case "business":
		user.VisBusiness += inc
	case "entertainment":
		user.VisEntertainment += inc
	case "general":
		user.VisGeneral += inc
	case "health":
		user.VisEntertainment += inc
	case "science":
		user.VisScience += inc
	case "sports":
		user.VisSports += inc
	case "technology":
		user.VisTechnology += inc
	default:
		return errors.New("news's category invalid")
	}

	return nil
}
