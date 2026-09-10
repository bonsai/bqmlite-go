package bqmlite

import "testing"

func TestFingerprintStable(t *testing.T) {
	p := Plan{Engine: "mean", Dataset: Dataset{Name: "x", Target: "y", Rows: []Row{{"y": 1}, {"y": 3}}}}
	a, err := PlanFingerprint(p); if err != nil { t.Fatal(err) }
	b, err := PlanFingerprint(p); if err != nil { t.Fatal(err) }
	if a != b { t.Fatalf("fingerprint changed: %s != %s", a, b) }
}
