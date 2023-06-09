package usecase

import (
	"errors"
	"github.com/jesusEstaba/calculator/pkg/domain"
)

type LoginUseCase struct {
	userRepo   domain.UserRepository
	passwdRepo domain.PasswordRepository
	tokenRepo  domain.TokenRepository
}

func NewLoginUseCase(
	userRepo domain.UserRepository,
	passwdRepo domain.PasswordRepository,
	tokenRepo domain.TokenRepository,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo,
		passwdRepo,
		tokenRepo,
	}
}

func (uc *LoginUseCase) Login(username, password string) (*string, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials: user does not exist")
	}

	err = uc.passwdRepo.Compare(password, user.Password)
	if err != nil {
		return nil, errors.New("invalid credentials: passwords does not match")
	}

	token, err := uc.tokenRepo.Generate(user.ID.Hex())
	if err != nil {
		return nil, errors.New("we can not be able to generate the token")
	}

	return &token, nil
}
