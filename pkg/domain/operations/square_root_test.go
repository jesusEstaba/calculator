package operations

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveSquareRoot(t *testing.T) {
	// given
	operation := &SquareRoot{}

	calc := &domain.Calculation{
		OperandA: 4,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(2)), result.Result)
}

func TestNegativeSquareRoot(t *testing.T) {
	// given
	operation := &SquareRoot{}

	calc := &domain.Calculation{
		OperandA: -4,
	}

	// when
	_, err := operation.Calculate(calc)

	// then
	assert.Equal(t, "can not perform square root of a negative number", err.Error())
}

func TestZeroSquareRoot(t *testing.T) {
	// given
	operation := &SquareRoot{}

	calc := &domain.Calculation{
		OperandA: 0,
	}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(0)), result.Result)
}
