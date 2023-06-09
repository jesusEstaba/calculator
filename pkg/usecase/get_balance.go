package usecase

import "github.com/jesusEstaba/calculator/pkg/domain"

type GetBalanceUseCase struct {
	repository domain.UserRepository
}

func NewGetBalanceUseCase(
	repository domain.UserRepository,
) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		repository,
	}
}

func (uc *GetBalanceUseCase) Balance(userID string) (float64, error) {
	user, err := uc.repository.GetUser(userID)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
