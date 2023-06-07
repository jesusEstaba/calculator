package operations

import (
	"errors"
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
	"math"
)

type SquareRoot struct{}

func (o *SquareRoot) Calculate(c *domain.Calculation) (*domain.CalculationResult, error) {
	if c.OperandA < 0 {
		return nil, errors.New("can not perform square root of a negative number")
	}

	sqr := math.Sqrt(c.OperandA)

	return &domain.CalculationResult{
		Result: fmt.Sprintf("%f", sqr),
	}, nil
}
