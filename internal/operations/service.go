package operations

import (
	"errors"
)

type CalculationService interface {
	Add(a, b float64) float64
	Sub(a, b float64) float64
	Mul(a, b float64) float64
	Div(a, b float64) (float64, error)
}

type ToTie struct{}

func (t ToTie) Add(a, b float64) (result float64) {
	return a + b
}

func (t ToTie) Sub(a, b float64) (result float64) {
	return a - b
}

func (t ToTie) Mul(a, b float64) (result float64) {
	return a * b
}

func (t ToTie) Div(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, err
}
