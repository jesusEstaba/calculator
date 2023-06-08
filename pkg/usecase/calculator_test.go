package usecase

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
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
