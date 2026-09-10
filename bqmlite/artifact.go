package bqmlite

import (
	"encoding/json"
	"fmt"
	"os"
)

// EngineCapabilities describes what an engine can execute.
type EngineCapabilities struct { Name string `json:"name"`; Train bool `json:"train"`; Predict bool `json:"predict"`; Deterministic bool `json:"deterministic"` }

type capabilityProvider interface { Capabilities() EngineCapabilities }
type artifact struct { Engine string `json:"engine"`; Model json.RawMessage `json:"model"` }

func (MeanEngine) Capabilities() EngineCapabilities { return EngineCapabilities{Name:"mean",Train:true,Predict:true,Deterministic:true} }
func (LinearRegressionEngine) Capabilities() EngineCapabilities { return EngineCapabilities{Name:"linear_regression",Train:true,Predict:true,Deterministic:true} }
func (LogisticRegressionEngine) Capabilities() EngineCapabilities { return EngineCapabilities{Name:"logistic_regression",Train:true,Predict:true,Deterministic:true} }

// Capabilities returns metadata when the registered engine exposes it.
func (r *Registry) Capabilities(name string) (EngineCapabilities, bool) { e,ok:=r.Get(name);if !ok{return EngineCapabilities{},false};if p,ok:=e.(capabilityProvider);ok{return p.Capabilities(),true};return EngineCapabilities{Name:e.Name(),Train:true, Predict:true},true }

func SaveModel(path string, engine string, model Model) error {if model==nil{return fmt.Errorf("model is nil")};b,err:=json.Marshal(artifact{Engine:engine,Model:mustJSON(model)});if err!=nil{return fmt.Errorf("encode model: %w",err)};return os.WriteFile(path,append(b,'\n'),0644)}
func mustJSON(v any) json.RawMessage {b,_:=json.Marshal(v);return b}

// LoadModel restores a model artifact using the engine named in the artifact.
func LoadModel(path string, r *Registry) (Model,error) {if r==nil{return nil,fmt.Errorf("registry is nil")};b,err:=os.ReadFile(path);if err!=nil{return nil,err};var a artifact;if err=json.Unmarshal(b,&a);err!=nil{return nil,fmt.Errorf("decode model: %w",err)};switch a.Engine{case "mean":var m meanModel;if err=json.Unmarshal(a.Model,&m);err!=nil{return nil,err};return m,nil;case "linear_regression":var m linearModel;if err=json.Unmarshal(a.Model,&m);err!=nil{return nil,err};return m,nil;case "logistic_regression":var m logisticModel;if err=json.Unmarshal(a.Model,&m);err!=nil{return nil,err};return m,nil;default:return nil,fmt.Errorf("unsupported model engine %q",a.Engine)}}
