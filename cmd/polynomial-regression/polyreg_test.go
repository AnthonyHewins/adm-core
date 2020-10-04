package main

import (
	"os"
	"reflect"
	"testing"

	"fmt"
	"github.com/AnthonyHewins/adm-core/api"
)

const polyreg = "/polyreg"

func TestMain(m *testing.M) {
	os.Setenv("MAX_DEGREE", "5")
	os.Setenv("MAX_ELEMENTS", "100")

	os.Exit(m.Run())
}

func Test_polynomialRegression(t *testing.T) {
	invalidNegative1 := -1
	invalidZero := 0
	validOne := 1
	invalidSix := 6

	tooLong := make([]float64, maxElements+1)

	tests := []struct {
		name  string
		args  *polyRegReq
		want  *api.Success
		want1 error
	}{
		{
			"Degree can't be -1",
			&polyRegReq{X: []float64{}, Y: []float64{}, MaxDeg: invalidNegative1},
			nil,
			fmt.Errorf("maxDeg must satisfy 0 <= maxDeg <= 5"),
		},
		{
			"Degree can't be 0",
			&polyRegReq{X: []float64{}, Y: []float64{}, MaxDeg: invalidZero},
			nil,
			fmt.Errorf("maxDeg must satisfy 0 <= maxDeg <= 5"),
		},
		{
			"Degree can't be greater than 5",
			&polyRegReq{X: []float64{}, Y: []float64{}, MaxDeg: invalidSix},
			nil,
			fmt.Errorf("maxDeg must satisfy 0 <= maxDeg <= 5"),
		},
		{
			"Length can't be nothing",
			&polyRegReq{X: []float64{}, Y: []float64{}, MaxDeg: validOne},
			nil,
			fmt.Errorf("must have len(x)==len(y) && maxDeg < len(x, y) <= 100, got len(x) = 0 and len(y) = 0"),
		},
		{
			"Lengths must match",
			&polyRegReq{X: []float64{1}, Y: []float64{}, MaxDeg: validOne},
			nil,
			fmt.Errorf("must have len(x)==len(y) && maxDeg < len(x, y) <= 100, got len(x) = 1 and len(y) = 0"),
		},
		{
			"Lengths cannot exceed maxLength",
			&polyRegReq{X: tooLong, Y: []float64{}, MaxDeg: validOne},
			nil,
			fmt.Errorf("must have len(x)==len(y) && maxDeg < len(x, y) <= 100, got len(x) = 101 and len(y) = 0"),
		},
		{
			"Matrix cannot be singular",
			&polyRegReq{X: []float64{1, 1}, Y: []float64{1, 1}, MaxDeg: validOne},
			nil,
			fmt.Errorf("matrix singular or near-singular with condition number +Inf"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Handler(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("polynomialRegression()\ngot1 = '%v',\nwant = '%v'", got, tt.want)
			}
			if !reflect.DeepEqual(got1.Error(), tt.want1.Error()) {
				t.Errorf("polynomialRegression()\ngot1 = '%v',\nwant = '%v'", got1, tt.want1)
			}
		})
	}
}
