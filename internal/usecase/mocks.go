package usecase

import (
	"github.com/jesusEstaba/calculator/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockOperationRepo struct {
	mock.Mock
}

func (m *MockOperationRepo) GetOperationCost(operationName string) (float64, error) {
	args := m.Called(operationName)
	return args.Get(0).(float64), nil
}
func (m *MockOperationRepo) RecordOperation(record domain.Record) error {
	args := m.Called(record)
	return args.Error(0)
}

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUser(userID string) (domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(domain.User), nil
}

func (m *MockUserRepo) FindByUsername(userID string) (*domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*domain.User), nil
}

func (m *MockUserRepo) Save(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), nil
}

func (m *MockUserRepo) UpdateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

type MockRandomStringRepo struct {
	mock.Mock
}

func (m *MockRandomStringRepo) Generate() (string, error) {
	return "", nil
}