package usecase

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
)

type RegisterUserUseCase struct {
	userRepo   domain.UserRepository
	passwdRepo domain.PasswordRepository
}

func NewRegisterUserUseCase(
	userRepo domain.UserRepository,
	passwdRepo domain.PasswordRepository,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo,
		passwdRepo,
	}
}

func (uc *RegisterUserUseCase) RegisterUser(username, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, err := uc.passwdRepo.Generate(password)
	if err != nil {
		return nil, err
	}
	newUser := domain.User{
		Username: username,
		Password: hashedPassword,
		Balance:  10000,
		Status:   domain.UserStatusActive,
	}

	created, err := uc.userRepo.Save(newUser)
	if err != nil {
		return nil, err
	}

	return &created, nil
}
