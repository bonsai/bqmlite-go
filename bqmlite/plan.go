package bqmlite

import (
	"context"
	"fmt"
)

// Plan describes one deterministic training + prediction execution.
type Plan struct {
	Dataset Dataset `json:"dataset"`
	Engine  string  `json:"engine"`
	Rows    []Row   `json:"rows,omitempty"`
}

// Run executes a plan through the registry.
func Run(ctx context.Context, plan Plan, registry *Registry) (PredictionResult, error) {
	if registry == nil { return PredictionResult{}, fmt.Errorf("registry is nil") }
	engine, ok := registry.Get(plan.Engine)
	if !ok { return PredictionResult{}, fmt.Errorf("engine %q not found", plan.Engine) }
	model, err := engine.Train(ctx, plan.Dataset)
	if err != nil { return PredictionResult{}, err }
	rows := plan.Rows
	if len(rows) == 0 { rows = plan.Dataset.Rows }
	out := PredictionResult{ModelName: model.Name(), Results: make([]Prediction, 0, len(rows))}
	for _, row := range rows {
		if err := ctx.Err(); err != nil { return PredictionResult{}, err }
		out.Results = append(out.Results, model.Predict(row))
	}
	return out, nil
}
