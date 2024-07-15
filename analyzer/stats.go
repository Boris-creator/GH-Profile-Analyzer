package analyzer

import (
	"math"
	"slices"
)

func getAverage[N int | float64](values []N, mapper *(func(N) N)) float64 {
	m := func(v N) N { return v }
	if mapper != nil {
		m = *mapper
	}
	var sum N = 0
	for _, v := range values {
		sum += m(v)
	}
	return float64(sum) / float64(len(values))
}
func getDeviation[N int | float64](values []N) float64 {
	pow := func(v N) N { return N(math.Pow(float64(v), 2)) }
	dispersion := getAverage(values, &pow) - math.Pow(getAverage(values, nil), 2)
	return math.Sqrt(dispersion)
}
func topN[K comparable](count int, list map[K]int) []K {
	keys := make([]K, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	slices.SortFunc[[]K](keys, func(k1, k2 K) int { return list[k2] - list[k1] })

	resCount := count
	if resCount > len(keys) {
		resCount = len(keys)
	}
	return keys[:resCount]
}
