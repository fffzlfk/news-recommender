package similarity

func Sim(a, b map[string]float32) (weight float32) {
	for k, v := range a {
		weight += v * b[k]
	}
	return
}
