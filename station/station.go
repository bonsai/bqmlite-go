package station

import "sort"

// Candidate is domain input with precomputed feature scores in [0,1].
type Candidate struct {
	Name            string             `json:"name"`
	Features        map[string]float64 `json:"features"`
	Reason          string             `json:"reason,omitempty"`
}

type Result struct {
	Name        string             `json:"name"`
	Score       float64            `json:"score"`
	Probability float64            `json:"probability"`
	Features    map[string]float64 `json:"features"`
	Reason      string             `json:"reason,omitempty"`
	Model       string             `json:"model"`
}

type Weights map[string]float64

var DefaultWeights = Weights{
	"phonetic": 0.30, "mora_rhythm": 0.20, "rhyme_structure": 0.20,
	"semantic": 0.10, "context_association": 0.10, "integrated_fit": 0.10,
}

func Suggest(candidates []Candidate, weights Weights, limit int) []Result {
	if len(weights) == 0 { weights = DefaultWeights }
	results := make([]Result, 0, len(candidates))
	for _, c := range candidates {
		var score, total float64
		for feature, weight := range weights { if value, ok := c.Features[feature]; ok { score += value * weight; total += weight } }
		if total > 0 { score /= total }
		results = append(results, Result{Name: c.Name, Score: score, Probability: score, Features: c.Features, Reason: c.Reason, Model: "station-weighted-v1"})
	}
	sort.SliceStable(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	if limit > 0 && len(results) > limit { results = results[:limit] }
	return results
}
