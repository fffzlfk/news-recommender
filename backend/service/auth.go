package service

import (
	"news-api/database"
	"news-api/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterService(data map[string]string) (*models.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if res := database.DB.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func LoginService(data map[string]string) (*models.User, error) {
	var user models.User
	if res := database.DB.Where("email = ?", data["email"]).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func UserService(id string) (*models.User, error) {
	var user models.User
	if res := database.DB.Where("id = ?", id).First(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
