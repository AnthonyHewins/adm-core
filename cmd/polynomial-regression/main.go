package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/AnthonyHewins/adm-core/api"
	"github.com/AnthonyHewins/adm-core/util"
	"github.com/AnthonyHewins/polyfit"
)

var (
	maxDegree   = int(util.EnvInt("MAX_DEGREE", 5))
	maxElements = int(util.EnvInt("MAX_ELEMENTS", 100))
)

type polyRegReq struct {
	MaxDeg int       `form:"maxDeg" json:"maxDeg"`
	X      []float64 `form:"x"      json:"x"`
	Y      []float64 `form:"y"      json:"y"`
}

func Handler(req *polyRegReq) (*api.Success, error) {
	if maxDegree < req.MaxDeg || 0 >= req.MaxDeg {
		return nil, fmt.Errorf("maxDeg must satisfy 0 <= maxDeg <= %v", maxDegree)
	}

	n := len(req.X)
	m := len(req.Y)

	if n != m || n > maxElements || n <= req.MaxDeg {
		return nil, fmt.Errorf("must have len(x)==len(y) && maxDeg < len(x, y) <= %v, got len(x) = %v and len(y) = %v", maxElements, n, m)
	}

	coef, err := polyfit.PolynomialRegression(req.X, req.Y, int(req.MaxDeg))
	if err != nil {
		return nil, err
	}

	return api.Ok(coef), nil
}

func main() {
	lambda.Start(Handler)
}
