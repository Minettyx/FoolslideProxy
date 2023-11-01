package utils

import (
	"strings"

	"github.com/Minettyx/FoolslideProxy/pkg/types"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

type DistanceSorter struct {
	Data      []*types.SearchResult
	Distances []float64
}

// TODO: temporary solution, find better algoritm
func (t DistanceSorter) ComputeDistances(query string) {
	sim := metrics.NewLevenshtein()
	sim.CaseSensitive = false

	for i := range t.Data {
		t.Distances[i] = strutil.Similarity(query, t.Data[i].Title, sim)
		if strings.Contains(strings.ToLower(t.Data[i].Title), strings.ToLower(query)) {
			t.Distances[i] += 1
		}
	}

}

func (a DistanceSorter) Len() int {
	return len(a.Data)
}

func (a DistanceSorter) Swap(i, j int) {
	a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
	a.Distances[i], a.Distances[j] = a.Distances[j], a.Distances[i]
}

func (a DistanceSorter) Less(i, j int) bool {
	return a.Distances[i] > a.Distances[j]
}
