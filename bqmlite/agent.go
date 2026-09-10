package bqmlite

import (
	"context"
	"encoding/json"
	"fmt"
)

// AgentRequest is the stable JSON boundary for OpenCode/Hermes-style callers.
type AgentRequest struct {
	Plan Plan `json:"plan"`
}

type AgentResponse struct {
	Result      PredictionResult `json:"result"`
	Fingerprint string           `json:"fingerprint"`
	Error       string           `json:"error,omitempty"`
}

func ExecuteJSON(ctx context.Context, input []byte, registry *Registry) ([]byte, error) {
	var req AgentRequest
	if err := json.Unmarshal(input, &req); err != nil { return nil, fmt.Errorf("decode request: %w", err) }
	result, err := Run(ctx, req.Plan, registry)
	if err != nil { return json.Marshal(AgentResponse{Error: err.Error()}) }
	fp, err := ResultFingerprint(result); if err != nil { return nil, err }
	return json.Marshal(AgentResponse{Result: result, Fingerprint: fp})
}
