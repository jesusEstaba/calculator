package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
)

type Multiplication struct{}

func (o *Multiplication) Calculate(c *domain.Calculation) (*domain.CalculationResult, error) {
	mul := c.OperandA * c.OperandB

	return &domain.CalculationResult{
		Result: fmt.Sprintf("%f", mul),
	}, nil
}
