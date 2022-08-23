package similarity

import (
	"news-api/models"
	"news-api/rpc/client"
	"sort"
)

type Recommender struct {
	client     *client.Client
	recentNews map[uint]map[string]float32
}

func NewRecommender(recentnews []models.News) (*Recommender, func()) {
	client, closeFunc := client.NewClient()

	mp := make(map[uint]map[string]float32)
	for _, news := range recentnews {
		mp[news.ID] = client.GetKeywords(news.Title)
	}
	return &Recommender{
		client:     client,
		recentNews: mp,
	}, closeFunc
}

func titleMatchValue(mpA, mpB map[string]float32) (value float32) {
	for k, v := range mpA {
		value += v * mpB[k] * 100.0
	}
	return
}

func (r *Recommender) newsMatchValue(mother, news models.News, motherMap map[string]float32) (value float32) {
	if mother.Category != "general" && mother.Category == news.Category {
		value += 4.0
	}

	v := titleMatchValue(motherMap, r.recentNews[news.ID])
	value += v

	if news.Source != "" && mother.Source == news.Source {
		value += 1.5
	}

	if news.Author != "" && mother.Author == news.Author {
		value += 1.5
	}

	return
}

func (r *Recommender) SimOrderNews(mother models.News, news []models.News) []models.News {

	motherMap := r.client.GetKeywords(mother.Title)

	sort.Slice(news, func(i, j int) bool {
		return r.newsMatchValue(mother, news[i], motherMap) > r.newsMatchValue(mother, news[j], motherMap)
	})

	return news[:10]
}
