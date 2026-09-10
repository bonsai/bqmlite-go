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

Core flow:

```text
Dataset → Train → Model → Predict → PredictionResult
```

## Implemented

- Core domain contracts: `Dataset`, `Row`, `Model`, `Prediction`, `PredictionResult`, `Engine`
- Deterministic dependency-free `MeanEngine`
- `Plan` + `Run` execution boundary
- Engine `Registry` and `AutoEngine`
- Local JSON dataset/result persistence
- Minimal BQML-style `CREATE MODEL` parser/translator
- Machine-readable Agent JSON boundary for OpenCode/Hermes
- SHA-256 plan/result fingerprints for reproducibility
- Station domain adapter with configurable feature weights
- Deterministic unit tests
- Thin CLI: `bqmlite -input dataset.json [-engine mean] [-output result.json]`

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
go run ./cmd/bqmlite -input dataset.json
```

## Design boundaries

- **Core** knows nothing about SQL, CLI, Station, agents, or ML frameworks.
- **SQL** translates into a plan; it never calls an engine directly.
- **Engine** is replaceable through the `Engine` interface.
- **Storage** is an adapter and currently uses standard-library JSON.
- **Agent** receives/sends JSON and does not contain ML logic.
- **Station** is a domain adapter outside the core.

## SQL subset

The parser intentionally supports only a small, explicit subset of `CREATE MODEL`. Unsupported syntax fails rather than pretending to be BigQuery-compatible. This keeps the implementation honest and makes expansion testable.

## Reproducibility

A `Plan` and `PredictionResult` can be fingerprinted with SHA-256. Deterministic engines should produce identical results for identical inputs and configuration.

## Non-goals

- BigQuery compatibility in full
- TensorFlow/Keras reimplementation
- legacy Python API compatibility
- hiding domain logic inside Agent skills

## Status

**Implemented foundation / first complete vertical slice.** Further model families can be added without changing the core contract.
