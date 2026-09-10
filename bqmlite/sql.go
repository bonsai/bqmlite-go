package bqmlite

import (
	"fmt"
	"regexp"
	"strings"
)

// SQLStatement is the intentionally small SQL subset understood by BQMLite.
type SQLStatement struct {
	Operation string
	Model     string
	Dataset   string
	Target    string
}

var createModelRE = regexp.MustCompile(`(?is)^\s*CREATE\s+MODEL\s+([A-Za-z0-9_.-]+)\s+OPTIONS\s*\(\s*model_type\s*=\s*['\"]([^'\"]+)['\"]\s*\)\s+AS\s+SELECT\s+(.+?)\s+FROM\s+([A-Za-z0-9_.-]+)\s*;?\s*$`)

func ParseSQL(sql string) (SQLStatement, error) {
	m := createModelRE.FindStringSubmatch(sql)
	if len(m) != 5 { return SQLStatement{}, fmt.Errorf("unsupported SQL: expected CREATE MODEL ... AS SELECT ... FROM ...") }
	selectPart := strings.TrimSpace(m[3])
	if !strings.Contains(strings.ToUpper(selectPart), "TARGET") { return SQLStatement{}, fmt.Errorf("SELECT clause must identify target") }
	return SQLStatement{Operation: "CREATE MODEL", Model: m[1], Dataset: m[4], Target: strings.TrimSpace(m[3])}, nil
}

// TranslateSQL maps supported BQML syntax to a backend-neutral Plan.
func TranslateSQL(sql string, dataset Dataset) (Plan, error) {
	stmt, err := ParseSQL(sql); if err != nil { return Plan{}, err }
	engine := strings.ToLower(stmt.Model)
	if strings.Contains(engine, "linear") { engine = "mean" }
	if dataset.Name == "" { dataset.Name = stmt.Dataset }
	return Plan{Dataset: dataset, Engine: engine}, nil
}
