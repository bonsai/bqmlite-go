package bqmlite

import (
	"context"
	"fmt"
)

type Registry struct{ engines map[string]Engine }

func NewRegistry(engines ...Engine) *Registry {
	r := &Registry{engines: map[string]Engine{}}
	for _, e := range engines { if e != nil { r.engines[e.Name()] = e } }
	return r
}

func (r *Registry) Register(e Engine) error {
	if r == nil { return fmt.Errorf("registry is nil") }
	if e == nil { return fmt.Errorf("engine is nil") }
	if r.engines == nil { r.engines = map[string]Engine{} }
	if _, exists := r.engines[e.Name()]; exists { return fmt.Errorf("engine %q already registered", e.Name()) }
	r.engines[e.Name()] = e
	return nil
}

func (r *Registry) Get(name string) (Engine, bool) { if r == nil { return nil, false }; e, ok := r.engines[name]; return e, ok }

// AutoEngine chooses a registered engine from an explicit hint, then defaults to mean.
func (r *Registry) AutoEngine(ctx context.Context, dataset Dataset, hint string) (Engine, error) {
	if hint != "" { e, ok := r.Get(hint); if !ok { return nil, fmt.Errorf("engine %q not found", hint) }; return e, nil }
	if e, ok := r.Get("mean"); ok { return e, nil }
	return nil, fmt.Errorf("no suitable engine for dataset %q", dataset.Name)
}
