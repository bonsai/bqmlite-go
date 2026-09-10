package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bonsai/bqmlite-go/bqmlite"
)

func main() {
	input := flag.String("input", "", "dataset JSON or CSV file")
	engine := flag.String("engine", "mean", "engine: mean, linear_regression, logistic_regression")
	output := flag.String("output", "", "optional result JSON file")
	flag.Parse()
	if *input == "" { fmt.Fprintln(os.Stderr, "usage: bqmlite -input dataset.json|csv [-engine mean|linear_regression|logistic_regression] [-output result.json]"); os.Exit(2) }
	var dataset bqmlite.Dataset
	var err error
	if strings.EqualFold(filepath.Ext(*input), ".csv") { dataset,err=bqmlite.LoadCSV(*input) } else { dataset,err=bqmlite.LoadDataset(*input) }
	if err != nil { fatal(err) }
	registry := bqmlite.NewRegistry(bqmlite.MeanEngine{}, bqmlite.LinearRegressionEngine{}, bqmlite.LogisticRegressionEngine{})
	result, err := bqmlite.Run(context.Background(), bqmlite.Plan{Dataset: dataset, Engine: *engine}, registry); if err != nil { fatal(err) }
	if *output != "" { if err := bqmlite.SaveResult(*output, result); err != nil { fatal(err) }; return }
	enc := json.NewEncoder(os.Stdout); enc.SetIndent("", "  "); if err := enc.Encode(result); err != nil { fatal(err) }
}

func fatal(err error) { fmt.Fprintln(os.Stderr, "bqmlite:", err); os.Exit(1) }
