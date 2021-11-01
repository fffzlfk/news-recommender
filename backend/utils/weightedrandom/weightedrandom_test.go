package weightedrandom

import (
	"testing"
)

func TestDo(t *testing.T) {
	wrc, err := NewChooser(map[string]int{
		"business":      1,
		"entertainment": 1,
		"general":       1,
		"health":        1,
		"science":       1,
		"sports":        1,
		"technology":    1,
	})
	if err != nil {
		t.Error(err)
	}

	t.Log(wrc.Pick())
}
