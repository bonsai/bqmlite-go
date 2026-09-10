package bqmlite

import "errors"

// MeanEngine is the first deterministic vertical-slice engine.
// It predicts the mean of a numeric target and intentionally has no ML
// framework dependency. More engines can implement Engine later.
type MeanEngine struct{}

type meanModel struct {
	mean float64
}

func (MeanEngine) Train(dataset Dataset) (Model, error) {
	if dataset.Target == "" {
		return nil, errors.New("target is required")
	}
	if len(dataset.Rows) == 0 {
		return nil, errors.New("dataset is empty")
	}

	var sum float64
	for _, row := range dataset.Rows {
		value, ok := row[dataset.Target].(float64)
		if !ok {
			return nil, errors.New("target must contain float64 values")
		}
		sum += value
	}

	return meanModel{mean: sum / float64(len(dataset.Rows))}, nil
}

func (m meanModel) Name() string { return "mean" }

func (m meanModel) Predict(row map[string]any) Prediction {
	return Prediction{Value: m.mean, Probability: 1.0}
}
