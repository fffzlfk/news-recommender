package similarity

import (
	"news-api/models"
	"news-api/rpc/client"
	"sort"
)

type Recommend struct {
	mother    models.News
	motherMap map[string]float32
}

func NewRecommend(mother models.News) *Recommend {
	client.Dial()
	return &Recommend{
		mother:    mother,
		motherMap: client.GetKeywords(mother.Title),
	}
}

func titleMatchValue(mpA, mpB map[string]float32) (value float32) {
	for k, v := range mpA {
		value += v * mpB[k]
	}
	return
}

func (r *Recommend) newsMatchValue(news models.News) (value float32) {
	if r.mother.Category == news.Category {
		value += 5.0
	}

	value += titleMatchValue(r.motherMap, client.GetKeywords(news.Title))

	if r.mother.Author == news.Author {
		value += 3.0
	}

	return
}

func (r *Recommend) SimOrderNews(news []models.News) []models.News {

	sort.Slice(news, func(i, j int) bool {
		return r.newsMatchValue(news[i]) > r.newsMatchValue(news[j])
	})

	client.Close()

	return news[:10]
}
