package models

import (
	"news-api/utils/weightedrandom"
)

type User struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email" gorm:"unique"`
	Password         []byte `json:"-"`
	VisBusiness      uint   `json:"vis_business" gorm:"default:1"`
	VisEntertainment uint   `json:"vis_entertainment" gorm:"default:1"`
	VisGeneral       uint   `json:"vis_general" gorm:"default:5"`
	VisHealth        uint   `json:"vis_health" gorm:"default:1"`
	VisScience       uint   `json:"vis_scientce" gorm:"default:1"`
	VisSports        uint   `json:"vis_sports" gorm:"default:1"`
	VisTechnology    uint   `json:"vis_technology" gorm:"default:1"`
}

func (user *User) GetANewsCategory() (string, error) {
	wrc, err := weightedrandom.NewChooser(map[string]uint{
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
