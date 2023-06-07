package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveMultiplication(t *testing.T) {
	// given
	operation := &Multiplication{}

	calc := &domain.Calculation{
		OperandA: 2,
		OperandB: 3,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(6)), result.Result)
}

func TestNegativeMultiplication(t *testing.T) {
	// given
	operation := &Multiplication{}

	calc := &domain.Calculation{
		OperandA: -2,
		OperandB: -3,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(6)), result.Result)
}

func TestZeroMultiplication(t *testing.T) {
	// given
	operation := &Multiplication{}

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
