package similarity

import (
	"news-api/models"
	"news-api/rpc/client"
	"sort"
)

type Recommender struct {
	recentNews map[uint]map[string]float32
}

func NewRecommender(recentnews []models.News) *Recommender {
	client.Dial()

	mp := make(map[uint]map[string]float32)
	for _, news := range recentnews {
		mp[news.ID] = client.GetKeywords(news.Title)
	}
	return &Recommender{
		recentNews: mp,
	}
}

func titleMatchValue(mpA, mpB map[string]float32) (value float32) {
	for k, v := range mpA {
		value += v * mpB[k]
	}
	return
}

func (r *Recommender) newsMatchValue(mother, news models.News, motherMap map[string]float32) (value float32) {
	if mother.Category == news.Category {
		value += 5.0
	}

	value += titleMatchValue(motherMap, r.recentNews[news.ID])

	if mother.Source == news.Source {
		value += 1.5
	}

	if mother.Author == news.Author {
		value += 1.5
	}

	return
}

func (r *Recommender) SimOrderNews(mother models.News, news []models.News) []models.News {

	motherMap := client.GetKeywords(mother.Title)

	sort.Slice(news, func(i, j int) bool {
		return r.newsMatchValue(mother, news[i], motherMap) > r.newsMatchValue(mother, news[j], motherMap)
	})

	return news[:10]
}

func (r *Recommender) Close() {
	client.Close()
}
