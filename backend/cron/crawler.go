package cron

import (
	"log"
	"net/http"
	"news-api/config"
	"news-api/database"
	"news-api/models"
	"sync"
	"time"

	"github.com/robtec/newsapi/api"
)

const apiKey = "d1f4a490e88f467ab34c822443df7b32"

func addNews() {
	var cnt int64
	database.DB.Model(models.News{}).Count(&cnt)
	if cnt < config.GetMaxNewsNumofDB() {
		crawlNews()
	}
}

func crawlNews() {
	client, err := api.New(&http.Client{}, apiKey, "https://newsapi.org")
	if err != nil {
		log.Println(err)
	}

	var wg sync.WaitGroup

	for _, category := range models.Categorys {
		wg.Add(1)
		go func(category string) {
			opts := api.Options{
				Country:  "cn",
				Category: category,
				PageSize: 100,
			}
			topHeadlines, err := client.TopHeadlines(opts)
			if err != nil {
				return
			}

			for _, art := range topHeadlines.Articles {

				news := models.News{
					Title:       art.Title,
					Description: art.Description,
					Url:         art.URL,
					UrlToImage:  art.URLToImage,
					Category:    category,
					Source:      art.Source.ID,
					Author:      art.Author,
					CreatedAt:   time.Now().Unix(),
				}
				database.DB.Create(&news)
			}
			wg.Done()
		}(category)

	}
	wg.Wait()
}
