package main

func average(values []float64) float64 {
	var total float64
	total = 0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}
