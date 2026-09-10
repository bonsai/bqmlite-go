package bqmlite

// Dataset is the input to a BQMLite training operation.
type Dataset struct {
	Name   string
	Rows   []map[string]any
	Target string
}

// Model is a trained model with a stable prediction contract.
type Model interface {
	Name() string
	Predict(row map[string]any) Prediction
}

// Prediction is one model output.
type Prediction struct {
	Value      any
	Probability float64
}

// PredictionResult contains predictions and model metadata.
type PredictionResult struct {
	ModelName string
	Results   []Prediction
}

// Engine trains a dataset into a model.
type Engine interface {
	Train(dataset Dataset) (Model, error)
}
