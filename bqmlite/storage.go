package bqmlite

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadDataset(path string) (Dataset, error) {
	b, err := os.ReadFile(path); if err != nil { return Dataset{}, err }
	var d Dataset
	if err := json.Unmarshal(b, &d); err != nil { return Dataset{}, fmt.Errorf("decode dataset: %w", err) }
	if d.Name == "" { d.Name = path }
	return d, nil
}

func SaveDataset(path string, d Dataset) error {
	b, err := json.MarshalIndent(d, "", "  "); if err != nil { return err }
	return os.WriteFile(path, append(b, '\n'), 0644)
}

func SaveResult(path string, r PredictionResult) error {
	b, err := json.MarshalIndent(r, "", "  "); if err != nil { return err }
	return os.WriteFile(path, append(b, '\n'), 0644)
}
