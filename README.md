# BQMLite Go

Full-scratch Go implementation of **BQMLite**: a small, deterministic local execution engine inspired by the BQML workflow, without depending on BigQuery or an ML framework.

> The legacy Python implementation is reference material only. It is not ported into this repository.

## Architecture

```text
SQL / JSON / Agent
        ↓
      Parser
        ↓
   Execution Plan
        ↓
      Registry
        ↓
      Engine
        ↓
 Model / Prediction
        ↓
      Result
        ↓
  Fingerprint
```

Core flow: `Dataset → Train → Model → Predict → PredictionResult`

## Implemented

- Core contracts: `Dataset`, `Row`, `Model`, `Prediction`, `PredictionResult`, `Engine`
- Deterministic dependency-free engines: `mean`, `linear_regression`, `logistic_regression`
- `Plan` + `Run` execution boundary
- Engine `Registry`, `AutoEngine`, and capability metadata
- JSON dataset/result persistence
- CSV dataset ingestion and schema validation
- JSON model artifact save/load
- Minimal BQML-style `CREATE MODEL` parser/translator
- Machine-readable Agent JSON boundary for OpenCode/Hermes
- SHA-256 plan/result fingerprints for reproducibility
- Station domain adapter with configurable feature weights
- Deterministic unit and regression/artifact tests
- Thin CLI with JSON/CSV input and selectable engines

## Dataset example

```json
{
  "name": "example",
  "target": "y",
  "rows": [
    {"x": 1, "y": 2},
    {"x": 2, "y": 4}
  ]
}
```

Run:

```bash
go test ./...
go vet ./...
go run ./cmd/bqmlite -input dataset.json -engine mean
go run ./cmd/bqmlite -input dataset.csv -engine linear_regression
```

## Design boundaries

- **Core** knows nothing about SQL, CLI, Station, agents, or ML frameworks.
- **SQL** translates into a plan; it never calls an engine directly.
- **Engine** is replaceable through the `Engine` interface.
- **Storage** is an adapter using standard-library JSON/CSV.
- **Agent** receives/sends JSON and does not contain ML logic.
- **Station** is a domain adapter outside the core.

The current implementation intentionally keeps the public package compact. Splitting every concern into separate packages is a refactoring option, not a prerequisite for the execution contract.

## SQL subset

The parser intentionally supports only a small, explicit subset of `CREATE MODEL`. Unsupported syntax fails rather than pretending to be BigQuery-compatible. This keeps the implementation honest and makes expansion testable.

## Reproducibility

A `Plan` and `PredictionResult` can be fingerprinted with SHA-256. The included engines are deterministic for identical inputs and configuration; regression engines infer feature order lexicographically and use fixed training parameters.

## Completion definition

BQMLite-Go is considered complete at the **local execution-core** level when a defined BQML subset can execute reproducibly as:

`Request → Plan → Engine → Model → Result`

It is **not** intended to become a full BigQuery ML compatibility layer.

## Non-goals

- BigQuery compatibility in full
- TensorFlow/Keras reimplementation
- legacy Python API compatibility
- hiding domain logic inside Agent skills
