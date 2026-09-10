package bqmlite

import (
	"context"
	"fmt"
)

// MeanEngine is a deterministic baseline engine with no ML dependency.
type MeanEngine struct{}

type meanModel struct{ mean float64 }

func (MeanEngine) Name() string { return "mean" }

func (MeanEngine) Train(ctx context.Context, dataset Dataset) (Model, error) {
	if err := ctx.Err(); err != nil { return nil, err }
	if dataset.Target == "" { return nil, fmt.Errorf("target is required") }
	if len(dataset.Rows) == 0 { return nil, fmt.Errorf("dataset is empty") }
	var sum float64
	for _, row := range dataset.Rows {
		if err := ctx.Err(); err != nil { return nil, err }
		v, ok := number(row[dataset.Target])
		if !ok { return nil, fmt.Errorf("target %q must contain numeric values", dataset.Target) }
		sum += v
	}
	return meanModel{mean: sum / float64(len(dataset.Rows))}, nil
}

func (m meanModel) Name() string { return "mean" }
func (m meanModel) Predict(Row) Prediction { return Prediction{Value: m.mean, Probability: 1} }

func number(v any) (float64, bool) {
	switch x := v.(type) {
	case float64: return x, true
	case float32: return float64(x), true
	case int: return float64(x), true
	case int8: return float64(x), true
	case int16: return float64(x), true
	case int32: return float64(x), true
	case int64: return float64(x), true
	case uint: return float64(x), true
	case uint8: return float64(x), true
	case uint16: return float64(x), true
	case uint32: return float64(x), true
	case uint64: return float64(x), true
	default: return 0, false
	}
}
