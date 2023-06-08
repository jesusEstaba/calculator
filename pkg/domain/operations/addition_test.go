package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveAddition(t *testing.T) {
	// given
	operation := &Addition{}

	calc := &domain.Calculation{
		OperandA: 2,
		OperandB: 3,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(5)), result.Result)
}

func TestNegativeAddition(t *testing.T) {
	// given
	operation := &Addition{}

	calc := &domain.Calculation{
		OperandA: -2,
		OperandB: -3,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(-5)), result.Result)
}

func TestZeroAddition(t *testing.T) {
	// given
	operation := &Addition{}

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
