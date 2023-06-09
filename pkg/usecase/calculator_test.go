package usecase

import (
	"fmt"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// given
	userID, _ := primitive.ObjectIDFromHex("648262f623eeafdfb68110e0")
	user := domain.User{
		ID:      &userID,
		Balance: 2000,
	}

	calculation := &domain.Calculation{
		OperationName: "addition",
		OperandA:      2,
		OperandB:      2,
	}

	operationID, _ := primitive.ObjectIDFromHex("operation-1")
	operation := &domain.Operation{
		ID:   &operationID,
		Cost: 1000,
	}
	
	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperation", calculation.OperationName).Return(operation, nil)
	operationRepo.On("RecordOperation", mock.MatchedBy(func(recordInput domain.Record) bool {
		return recordInput.Amount == 1000 && recordInput.UserBalance == 1000
	})).Return(nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", userID.Hex()).Return(&user)

	updatedUser := domain.User{
		ID:      &userID,
		Balance: 1000,
	}
	userRepo.On("UpdateUser", updatedUser).Return(nil)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	result, err := usecase.Calculate(userID.Hex(), calculation)

	// then
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%f", float64(4)), result.Result)

	userRepo.AssertNumberOfCalls(t, "UpdateUser", 1)
	operationRepo.AssertNumberOfCalls(t, "RecordOperation", 1)
}

func TestCalculateWhenOperationFails(t *testing.T) {
	// given
	userID, _ := primitive.ObjectIDFromHex("648262f623eeafdfb68110e0")
	user := domain.User{
		ID:      &userID,
		Balance: 2000,
	}

	calculation := &domain.Calculation{
		OperationName: "division",
		OperandA:      2,
		OperandB:      0,
	}

	operationID, _ := primitive.ObjectIDFromHex("operation-1")
	operation := &domain.Operation{
		ID:   &operationID,
		Cost: 1000,
	}

	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperation", calculation.OperationName).Return(operation, nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", userID.Hex()).Return(&user)

	updatedUser := domain.User{
		ID:      &userID,
		Balance: 1000,
	}
	userRepo.On("UpdateUser", updatedUser).Return(nil)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	_, err := usecase.Calculate(userID.Hex(), calculation)

	// then
	assert.Equal(t, "can not perform division by zero", err.Error())
	userRepo.AssertNotCalled(t, "UpdateUser")
	operationRepo.AssertNotCalled(t, "RecordOperation")
}

func TestCalculateWithInsufficientBalance(t *testing.T) {
	// given
	userID, _ := primitive.ObjectIDFromHex("648262f623eeafdfb68110e0")
	user := domain.User{
		ID:      &userID,
		Balance: 0,
	}

	calculation := &domain.Calculation{
		OperationName: "addition",
		OperandA:      2,
		OperandB:      2,
	}

	operationID, _ := primitive.ObjectIDFromHex("operation-1")
	operation := &domain.Operation{
		ID:   &operationID,
		Cost: 1000,
	}

	operationRepo := new(MockOperationRepo)
	operationRepo.On("GetOperation", calculation.OperationName).Return(operation, nil)

	userRepo := new(MockUserRepo)
	userRepo.On("GetUser", userID.Hex()).Return(&user)

	randomRepo := new(MockRandomStringRepo)

	usecase := NewCalculatorUseCase(
		userRepo,
		operationRepo,
		randomRepo,
	)

	// when
	_, err := usecase.Calculate(userID.Hex(), calculation)

	// then
	assert.Equal(t, "insufficient balance", err.Error())
	userRepo.AssertNotCalled(t, "UpdateUser")
	operationRepo.AssertNotCalled(t, "RecordOperation")
}
