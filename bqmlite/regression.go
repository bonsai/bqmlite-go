package bqmlite

import (
	"context"
	"fmt"
	"math"
	"sort"
)

// LinearRegressionEngine fits ordinary least squares with a deterministic
// gradient-free normal-equation solver. Numeric feature columns are inferred
// from the first row and sorted by name for stable model artifacts.
type LinearRegressionEngine struct{}

type linearModel struct {
	Features []string `json:"features"`
	Weights  []float64 `json:"weights"`
	Bias     float64 `json:"bias"`
}

func (LinearRegressionEngine) Name() string { return "linear_regression" }
func (LinearRegressionEngine) Train(ctx context.Context, d Dataset) (Model, error) {
	if err := validateRegressionDataset(d); err != nil { return nil, err }
	features, X, y, err := numericMatrix(d)
	if err != nil { return nil, err }
	if err := ctx.Err(); err != nil { return nil, err }
	// Normal equations with a tiny diagonal regularizer make singular fixture
	// datasets deterministic and avoid external linear algebra dependencies.
	n := len(features) + 1
	A := make([][]float64, n); b := make([]float64, n)
	for i := range A { A[i] = make([]float64, n) }
	for r := range X {
		v := append([]float64{1}, X[r]...)
		for i := 0; i < n; i++ { b[i] += v[i] * y[r]; for j := 0; j < n; j++ { A[i][j] += v[i] * v[j] } }
	}
	for i := 0; i < n; i++ { A[i][i] += 1e-10 }
	w, err := solve(A, b); if err != nil { return nil, err }
	return linearModel{Features: features, Bias: w[0], Weights: w[1:]}, nil
}
func (m linearModel) Name() string { return "linear_regression" }
func (m linearModel) Predict(r Row) Prediction { v := m.Bias; for i, f := range m.Features { x, _ := number(r[f]); v += x*m.Weights[i] }; return Prediction{Value:v} }

// LogisticRegressionEngine uses deterministic batch gradient descent.
type LogisticRegressionEngine struct{}
type logisticModel struct { Features []string `json:"features"`; Weights []float64 `json:"weights"`; Bias float64 `json:"bias"` }
func (LogisticRegressionEngine) Name() string { return "logistic_regression" }
func (LogisticRegressionEngine) Train(ctx context.Context, d Dataset) (Model, error) {
	if err := validateRegressionDataset(d); err != nil { return nil, err }
	features, X, y, err := numericMatrix(d); if err != nil { return nil, err }
	for i, v := range y { if v != 0 && v != 1 { return nil, fmt.Errorf("target %q must contain only 0 or 1", d.Target) }; y[i]=v }
	w:=make([]float64,len(features)); b:=0.0; lr:=0.1
	for epoch:=0; epoch<1000; epoch++ { if err:=ctx.Err(); err!=nil{return nil,err}; gw:=make([]float64,len(w)); gb:=0.0; for r:=range X { z:=b; for j,x:=range X[r]{z+=x*w[j]}; p:=sigmoid(z); e:=p-y[r]; gb+=e; for j,x:=range X[r]{gw[j]+=e*x} }; scale:=1/float64(len(X)); b-=lr*gb*scale; for j:=range w{w[j]-=lr*gw[j]*scale} }
	return logisticModel{Features:features,Weights:w,Bias:b},nil
}
func (m logisticModel) Name() string{return "logistic_regression"}
func (m logisticModel) Predict(r Row) Prediction {z:=m.Bias;for i,f:=range m.Features{x,_:=number(r[f]);z+=x*m.Weights[i]};p:=sigmoid(z);v:=0;if p>=0.5{v=1};return Prediction{Value:v,Probability:p}}
func sigmoid(x float64) float64 { if x>=0 {e:=math.Exp(-x);return 1/(1+e)};e:=math.Exp(x);return e/(1+e) }
func validateRegressionDataset(d Dataset) error {if d.Target==""{return fmt.Errorf("target is required")};if len(d.Rows)==0{return fmt.Errorf("dataset is empty")};return nil}
func numericMatrix(d Dataset)([]string,[][]float64,[]float64,error){keys:=map[string]bool{};for _,r:=range d.Rows{for k:=range r{if k!=d.Target{keys[k]=true}}};features:=make([]string,0,len(keys));for k:=range keys{features=append(features,k)};sort.Strings(features);X:=make([][]float64,len(d.Rows));y:=make([]float64,len(d.Rows));for i,r:=range d.Rows{var ok bool;y[i],ok=number(r[d.Target]);if !ok{return nil,nil,nil,fmt.Errorf("target %q must be numeric",d.Target)};X[i]=make([]float64,len(features));for j,f:=range features{X[i][j],ok=number(r[f]);if !ok{return nil,nil,nil,fmt.Errorf("feature %q must be numeric",f)}}};return features,X,y,nil}
func solve(A [][]float64,b []float64)([]float64,error){n:=len(b);for i:=0;i<n;i++{p:=i;for r:=i+1;r<n;r++{if math.Abs(A[r][i])>math.Abs(A[p][i]){p=r}};if math.Abs(A[p][i])<1e-14{return nil,fmt.Errorf("singular matrix")};A[i],A[p]=A[p],A[i];b[i],b[p]=b[p],b[i];for r:=i+1;r<n;r++{q:=A[r][i]/A[i][i];for c:=i;c<n;c++{A[r][c]-=q*A[i][c]};b[r]-=q*b[i]}};x:=make([]float64,n);for i:=n-1;i>=0;i--{s:=b[i];for c:=i+1;c<n;c++{s-=A[i][c]*x[c]};x[i]=s/A[i][i]};return x,nil}
