package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
)

type Addition struct{}

func (o *Addition) Calculate(c *domain.Calculation) (*domain.CalculationResult, error) {
	sum := c.OperandA + c.OperandB

	return &domain.CalculationResult{
		Result: fmt.Sprintf("%f", sum),
	}, nil
}
