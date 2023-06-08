package usecase

import (
	"fmt"
	"github.com/jesusEstaba/calculator/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculate(t *testing.T) {
	// given
	user := domain.User{
		ID:      "user-1",
		Balance: 1000,
	}

	operation := &domain.Calculation{
		OperationName: "addition",
		OperandA:      2,
		OperandB:      2,
	}

	record := domain.Record{
		OperationID: "addition",
		UserID:      "user-1",
		Amount:      1000,
		UserBalance: 0,
		OperationResponse: &domain.CalculationResult{
			Result: fmt.Sprintf("%f", float64(4)),
		},
	}

	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperationCost", operation.OperationName).Return(float64(1000), nil)
	operationRepo.On("RecordOperation", record).Return(nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", user.ID).Return(user)

	user.Balance = 0
	userRepo.On("UpdateUser", user).Return(nil)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	result, err := usecase.Calculate(user.ID, operation)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(4)), result.Result)

	userRepo.AssertNumberOfCalls(t, "UpdateUser", 1)
	operationRepo.AssertNumberOfCalls(t, "RecordOperation", 1)
}

func TestCalculateWhenOperationFails(t *testing.T) {
	// given
	user := domain.User{
		ID:      "user-1",
		Balance: 1000,
	}

	operation := &domain.Calculation{
		OperationName: "division",
		OperandA:      2,
		OperandB:      0,
	}

	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperationCost", operation.OperationName).Return(float64(1000), nil)
	operationRepo.On("RecordOperation", domain.Record{}).Return(nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", user.ID).Return(user)
	userRepo.On("UpdateUser", user).Return(nil)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	_, err := usecase.Calculate(user.ID, operation)

	// then
	assert.Equal(t, "can not perform division by zero", err.Error())
	userRepo.AssertNotCalled(t, "UpdateUser")
	operationRepo.AssertNotCalled(t, "RecordOperation")
}

func TestCalculateWithInsufficientBalance(t *testing.T) {
	// given
	user := domain.User{
		ID:      "user-1",
		Balance: 0,
	}

	operation := &domain.Calculation{
		OperationName: "addition",
		OperandA:      2,
		OperandB:      2,
	}

	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperationCost", operation.OperationName).Return(float64(1000), nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", user.ID).Return(user)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	_, err := usecase.Calculate(user.ID, operation)

	// then
	assert.Equal(t, "insufficient balance", err.Error())
	userRepo.AssertNotCalled(t, "UpdateUser")
	operationRepo.AssertNotCalled(t, "RecordOperation")
}

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
