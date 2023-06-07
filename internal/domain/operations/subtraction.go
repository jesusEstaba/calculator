package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
)

type Subtraction struct{}

func (o *Subtraction) Calculate(c *domain.Calculation) (*domain.CalculationResult, error) {
	sub := c.OperandA - c.OperandB

	return &domain.CalculationResult{
		Result: fmt.Sprintf("%f", sub),
	}, nil
}
