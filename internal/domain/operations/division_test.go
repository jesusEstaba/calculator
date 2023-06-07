package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveDivision(t *testing.T) {
	// given
	operation := &Division{}

	calc := &domain.Calculation{
		OperandA: 6,
		OperandB: 2,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(3)), result.Result)
}

func TestNegativeDivision(t *testing.T) {
	// given
	operation := &Division{}

	calc := &domain.Calculation{
		OperandA: -6,
		OperandB: -2,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(3)), result.Result)
}

func TestZeroDividendDivision(t *testing.T) {
	// given
	operation := &Division{}

	calc := &domain.Calculation{
		OperandA: 0,
		OperandB: 2,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(0)), result.Result)
}

func TestZeroDivisorDivision(t *testing.T) {
	// given
	operation := &Division{}

	calc := &domain.Calculation{
		OperandA: 2,
		OperandB: 0,
	}

	// when
	_, err := operation.Calculate(calc)

	// then
	assert.Equal(t, "can not perform division by zero", err.Error())
}
