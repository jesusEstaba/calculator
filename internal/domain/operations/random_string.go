package operations

import (
	"errors"
	"github.com/jesusEstaba/calculator/internal/domain"
)

type RandomString struct {
	RandomRepo domain.RandomStringRepository
}

func (o *RandomString) Calculate(*domain.Calculation) (*domain.CalculationResult, error) {
	random, err := o.RandomRepo.Generate()
	if err != nil {
		return nil, errors.New("can not perform random string generation")
	}

	return &domain.CalculationResult{Result: random}, nil
}
