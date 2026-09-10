package bqmlite

import (
	"context"
	"testing"
)

func TestMeanEngine(t *testing.T) {
	dataset := Dataset{Name: "example", Target: "y", Rows: []Row{{"y": 2}, {"y": 4}}}
	model, err := (MeanEngine{}).Train(context.Background(), dataset)
	if err != nil { t.Fatal(err) }
	prediction := model.Predict(Row{})
	if prediction.Value != float64(3) { t.Fatalf("got %v, want 3", prediction.Value) }
}

func TestRun(t *testing.T) {
	dataset := Dataset{Name: "example", Target: "y", Rows: []Row{{"y": 2}, {"y": 4}}}
	result, err := Run(context.Background(), Plan{Dataset: dataset, Engine: "mean"}, NewRegistry(MeanEngine{}))
	if err != nil { t.Fatal(err) }
	if result.ModelName != "mean" || len(result.Results) != 2 { t.Fatalf("unexpected result: %+v", result) }
}

func TestRegistry(t *testing.T) {
	r := NewRegistry(MeanEngine{})
	if _, ok := r.Get("mean"); !ok { t.Fatal("mean engine not registered") }
	if _, err := r.AutoEngine(context.Background(), Dataset{}, "mean"); err != nil { t.Fatal(err) }
}
