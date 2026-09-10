package bqmlite

import (
	"context"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestLinearRegression(t *testing.T) {
	d:=Dataset{Target:"y",Rows:[]Row{{"x":1.0,"y":3.0},{"x":2.0,"y":5.0},{"x":3.0,"y":7.0}}}
	m,err:=(LinearRegressionEngine{}).Train(context.Background(),d);if err!=nil{t.Fatal(err)}
	p:=m.Predict(Row{"x":4.0});v,ok:=p.Value.(float64);if !ok||math.Abs(v-9)>1e-5{t.Fatalf("prediction=%v",p.Value)}
}
func TestLogisticRegression(t *testing.T) {
	d:=Dataset{Target:"y",Rows:[]Row{{"x":-2.0,"y":0},{"x":-1.0,"y":0},{"x":1.0,"y":1},{"x":2.0,"y":1}}}
	m,err:=(LogisticRegressionEngine{}).Train(context.Background(),d);if err!=nil{t.Fatal(err)}
	if p:=m.Predict(Row{"x":2.0});p.Value!=1||p.Probability<=0.5{t.Fatalf("prediction=%v",p)}
}
func TestModelArtifactRoundTrip(t *testing.T) {
	d:=Dataset{Target:"y",Rows:[]Row{{"x":1.0,"y":3.0},{"x":2.0,"y":5.0}}};m,err:=(LinearRegressionEngine{}).Train(context.Background(),d);if err!=nil{t.Fatal(err)}
	r:=NewRegistry(LinearRegressionEngine{});p:=filepath.Join(t.TempDir(),"model.json");if err=SaveModel(p,"linear_regression",m);err!=nil{t.Fatal(err)};got,err:=LoadModel(p,r);if err!=nil{t.Fatal(err)};a:=got.Predict(Row{"x":3.0});if a.Value==nil{t.Fatal("empty prediction")};_ = os.Remove(p)
}
