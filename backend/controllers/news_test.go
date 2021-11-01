package controllers

import (
	"fmt"
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"testing"
)

func TestLikeStateHandler(t *testing.T) {
	config.Init()
	database.Connect()
	var user models.User
	user.Id = 5
	database.DB.Model(&user).Association("LikedNews").Find(&user.LikedNews, []int{2245})
	fmt.Println(user.LikedNews)
	if len(user.LikedNews) != 1 {
		t.Fail()
	}
}

func TestLikeCountHandler(t *testing.T) {
	config.Init()
	database.Connect()
	var news models.News
	news.Id = 2245
	cnt := database.DB.Model(&news).Association("BeLikedBy").Count()
	fmt.Println(cnt)
	if cnt != 1 {
		t.Fail()
	}
}
