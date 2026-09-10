package station

import "testing"

func TestSuggest(t *testing.T) {
	got := Suggest([]Candidate{
		{Name: "五反田", Features: map[string]float64{"phonetic": .9, "mora_rhythm": .8}},
		{Name: "新宿", Features: map[string]float64{"phonetic": .4, "mora_rhythm": .4}},
	}, DefaultWeights, 1)
	if len(got) != 1 || got[0].Name != "五反田" { t.Fatalf("unexpected ranking: %+v", got) }
}
