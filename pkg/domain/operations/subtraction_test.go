package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveSubtraction(t *testing.T) {
	// given
	operation := &Subtraction{}

	calc := &domain.Calculation{
		OperandA: 3,
		OperandB: 2,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(1)), result.Result)
}

func TestNegativeSubtraction(t *testing.T) {
	// given
	operation := &Subtraction{}

	calc := &domain.Calculation{
		OperandA: -2,
		OperandB: -3,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(1)), result.Result)
}

func TestZeroSubtraction(t *testing.T) {
	// given
	operation := &Subtraction{}

	calc := &domain.Calculation{
		OperandA: 0,
		OperandB: 0,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(0)), result.Result)
}
