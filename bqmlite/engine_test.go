package bqmlite

import "testing"

func TestMeanEngine(t *testing.T) {
	dataset := Dataset{
		Name:   "example",
		Target: "y",
		Rows: []map[string]any{
			{"y": float64(2)},
			{"y": float64(4)},
		},
	}

	model, err := (MeanEngine{}).Train(dataset)
	if err != nil {
		t.Fatal(err)
	}

	prediction := model.Predict(map[string]any{})
	if prediction.Value != float64(3) {
		t.Fatalf("got %v, want 3", prediction.Value)
	}
}
