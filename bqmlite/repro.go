package bqmlite

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func Fingerprint(v any) (string, error) {
	b, err := json.Marshal(v); if err != nil { return "", fmt.Errorf("fingerprint: %w", err) }
	h := sha256.Sum256(b)
	return fmt.Sprintf("sha256:%x", h[:]), nil
}

func PlanFingerprint(p Plan) (string, error) { return Fingerprint(p) }
func ResultFingerprint(r PredictionResult) (string, error) { return Fingerprint(r) }
