package bqmlite

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// LoadCSV reads a header-based CSV. Values are parsed as bool, integer,
// float, or string in that order; empty cells remain empty strings.
func LoadCSV(path string) (Dataset, error) {
	f, err := os.Open(path); if err != nil { return Dataset{}, err }; defer f.Close()
	cr:=csv.NewReader(f); header,err:=cr.Read();if err!=nil{return Dataset{},fmt.Errorf("read CSV header: %w",err)}
	rows:=make([]Row,0);for {rec,e:=cr.Read();if e==io.EOF{break};if e!=nil{return Dataset{},fmt.Errorf("read CSV row: %w",e)};if len(rec)!=len(header){return Dataset{},fmt.Errorf("CSV row has %d fields; want %d",len(rec),len(header))};r:=Row{};for i,k:=range header{r[k]=parseCSVValue(rec[i])};rows=append(rows,r)}
	return Dataset{Name:path,Rows:rows},nil
}
func parseCSVValue(s string) any {if s==""{return s};if v,e:=strconv.ParseBool(s);e==nil{return v};if v,e:=strconv.ParseInt(s,10,64);e==nil{return v};if v,e:=strconv.ParseFloat(s,64);e==nil{return v};return s}

// ValidateDataset checks target presence and consistent non-empty column names.
func ValidateDataset(d Dataset) error {if len(d.Rows)==0{return fmt.Errorf("dataset is empty")};seen:=map[string]bool{};for _,r:=range d.Rows{for k:=range r{if k==""{return fmt.Errorf("schema contains empty column name")};seen[k]=true}};if d.Target!=""{if !seen[d.Target]{return fmt.Errorf("target %q not found in dataset",d.Target)}};return nil}
