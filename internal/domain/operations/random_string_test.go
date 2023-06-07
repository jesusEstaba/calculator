package operations

import (
	"errors"
	"github.com/jesusEstaba/calculator/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessRandomStringGeneration(t *testing.T) {
	// given
	operation := &RandomString{
		RandomRepo: &FakeSuccessRandomStringGenerator{},
	}

	calc := &domain.Calculation{}

	// when
	result, err := operation.Calculate(calc)

	// then
	assert.Nil(t, err)
	assert.Equal(t, "r4nd0m", result.Result)
}

func TestFailedRandomStringGeneration(t *testing.T) {
	// given
	operation := &RandomString{
		RandomRepo: &FakeFailureRandomStringGenerator{},
	}

	calc := &domain.Calculation{}

	// when
	_, err := operation.Calculate(calc)

	// then
	assert.Equal(t, "can not perform random string generation", err.Error())
}

type FakeSuccessRandomStringGenerator struct{}

func (o *FakeSuccessRandomStringGenerator) Generate() (string, error) {
	return "r4nd0m", nil
}

type FakeFailureRandomStringGenerator struct{}

func (o *FakeFailureRandomStringGenerator) Generate() (string, error) {
	return "", errors.New("fail")
}
