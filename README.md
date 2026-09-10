# BQMLite Go

Full-scratch Go implementation of **BQMLite**, a local BQML-compatible execution engine.

> This repository is the new implementation. It does not migrate the legacy Python implementation into Go.

## Direction

```text
BQML specification
      ↓
SQL / API
      ↓
Execution Plan
      ↓
BQMLite Engine
      ↓
Model / Prediction / Result
```

The core contract is:

```text
Dataset → Train → Model → Predict → PredictionResult
```

## Current vertical slice

- Go module
- Core `Dataset`, `Model`, `Prediction`, `PredictionResult`, `Engine` contracts
- Deterministic `MeanEngine` as the first dependency-free engine
- Minimal `bqmlite` CLI
- Unit test for the first engine

## Planned layers

```text
cmd/bqmlite/
bqmlite/
  core/
  sql/
  plan/
  engine/
  storage/
  model/
  result/
  station/
examples/
experiments/
.opencode/skills/bqmlite/
.hermes/skills/bqmlite/
```

## Principles

1. Full scratch: design the required engine instead of cleaning the old implementation.
2. Go-first: native binary, clear interfaces, local execution.
3. Python is an adapter/experiment layer, not the core.
4. SQL is a front-end; execution is driven by an explicit plan.
5. Agent interfaces remain thin and call the stable engine contract.
6. Station suggestion is a domain application, not part of the ML core.

## Status

Early foundation. The first goal is a small, deterministic vertical slice before adding SQL, storage, model families, and station suggestion.
