package matrix

import "math"

type Matrix struct {
	Content [][]int
	Dim     int
}

func (m Matrix) Mean() float64 {
	total := 0
	for idx := range m.Content {
		for jdx := range m.Content[idx] {

			total += m.Content[idx][jdx]
		}
	}
	return float64(total) / float64(m.Dim)
}

func (m Matrix) Variance() float64 {

	mean := m.Mean()
	var sum float64 = 0
	for idx := range m.Content {
		for jdx := range m.Content[idx] {
			meanDiff := float64(m.Content[idx][jdx]) - mean
			sum += meanDiff * meanDiff
		}
	}
	return sum / float64(m.Dim)
}

func (m Matrix) Std() float64 {
	return math.Sqrt(m.Variance())
}
