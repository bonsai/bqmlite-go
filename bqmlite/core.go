package bqmlite

import "context"

// Row is one dataset record.
type Row map[string]any

// Dataset is immutable input data for an execution.
type Dataset struct {
	Name   string `json:"name"`
	Rows   []Row  `json:"rows"`
	Target string `json:"target,omitempty"`
}

// Model is a trained model with a stable prediction contract.
type Model interface {
	Name() string
	Predict(Row) Prediction
}

// Prediction is one model output.
type Prediction struct {
	Value       any     `json:"value"`
	Probability float64 `json:"probability,omitempty"`
}

// PredictionResult is the reproducible output of an execution.
type PredictionResult struct {
	ModelName string       `json:"model_name"`
	Results   []Prediction `json:"results"`
}

// Engine trains datasets and produces models.
type Engine interface {
	Name() string
	Train(context.Context, Dataset) (Model, error)
}
