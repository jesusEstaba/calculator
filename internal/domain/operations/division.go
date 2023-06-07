package operations

import (
	"errors"
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
)

type Division struct{}

func (o *Division) Calculate(c *domain.Calculation) (*domain.CalculationResult, error) {
	if c.OperandB == 0 {
		return nil, errors.New("can not perform division by zero")
	}

	div := c.OperandA / c.OperandB

	return &domain.CalculationResult{
		Result: fmt.Sprintf("%f", div),
	}, nil
}
