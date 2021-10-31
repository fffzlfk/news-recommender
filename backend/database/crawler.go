package database

import (
	"log"
	"net/http"
	"news-api/models"
	"time"

	"github.com/robtec/newsapi/api"
)

const apiKey = "d1f4a490e88f467ab34c822443df7b32"

func crawlNew() {
	client, err := api.New(&http.Client{}, apiKey, "https://newsapi.org")
	if err != nil {
		log.Println(err)
	}

	for _, category := range models.Categorys {
		opts := api.Options{
			Country:  "cn",
			Category: category,
			PageSize: 100,
		}
		topHeadlines, err := client.TopHeadlines(opts)
		if err != nil {
			continue
		}

		for _, art := range topHeadlines.Articles {

			news := models.News{
				Title:       art.Title,
				Description: art.Description,
				Url:         art.URL,
				UrlToImage:  art.URLToImage,
				Category:    category,
				CreatedAt:   time.Now().Unix(),
			}
			DB.Create(&news)
		}
	}
}
