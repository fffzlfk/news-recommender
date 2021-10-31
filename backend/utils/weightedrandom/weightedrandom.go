package weightedrandom

import (
	"errors"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Choice struct {
	Item   string
	Weight uint
}

func NewChoice(item string, weight uint) Choice {
	return Choice{Item: item, Weight: weight}
}

type Chooser struct {
	data   []Choice
	totals []int
	max    int
}

var errWeightOverflow = errors.New("sum of Choice Weigths extends max int")
var errNoValidChoices = errors.New("zero Choices with Weight >= 1")

func NewChooser(choicesMap map[string]uint) (*Chooser, error) {
	choices := make([]Choice, 0, len(choicesMap))

	for k, v := range choicesMap {
		choices = append(choices, Choice{
			Item:   k,
			Weight: v,
		})
	}

	sort.Slice(choices, func(i, j int) bool {
		return choices[i].Weight < choices[j].Weight
	})

	totals := make([]int, len(choices))
	runningTotal := 0

	for i, c := range choices {
		weight := int(c.Weight)

		if (math.MaxInt - runningTotal) <= weight {
			return nil, errWeightOverflow
		}
		runningTotal += weight
		totals[i] = runningTotal
	}

	if runningTotal < 0 {
		return nil, errNoValidChoices
	}

	return &Chooser{data: choices, totals: totals, max: runningTotal}, nil
}

func (c *Chooser) Pick() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(c.max) + 1
	i := searchInts(c.totals, r)
	return c.data[i].Item
}

func searchInts(a []int, x int) int {
	i, j := 0, len(a)
	for i < j {
		h := int(uint(i+j) >> 1) // 防止溢出
		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
