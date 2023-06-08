package usecase

import (
	"github.com/jesusEstaba/calculator/internal/domain"
	"github.com/jesusEstaba/calculator/internal/domain/operations"
)

type CalculatorUseCase struct {
	operationRepo domain.CalculatorRepository
	userRepo      domain.UserRepository
	operations    map[string]domain.Operation
}

func NewCalculatorUseCase(
	userRepo domain.UserRepository,
	operationRepo domain.CalculatorRepository,
	randomStringRepo domain.RandomStringRepository,
) *CalculatorUseCase {
	operationList := make(map[string]domain.Operation)
	operationList["addition"] = &operations.Addition{}
	operationList["subtraction"] = &operations.Subtraction{}
	operationList["multiplication"] = &operations.Multiplication{}
	operationList["division"] = &operations.Division{}
	operationList["square_root"] = &operations.SquareRoot{}
	operationList["random_string"] = &operations.RandomString{
		RandomRepo: randomStringRepo,
	}

	return &CalculatorUseCase{
		userRepo:      userRepo,
		operationRepo: operationRepo,
		operations:    operationList,
	}
}

func (uc *CalculatorUseCase) Calculate(userID string, operation *domain.Calculation) (*domain.CalculationResult, error) {
	cost, err := uc.operationRepo.GetOperationCost(operation.OperationName)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}

	err = user.Withdraw(cost)
	if err != nil {
		return nil, err
	}

	operationFunc := uc.operations[operation.OperationName]
	result, err := operationFunc.Calculate(operation)
	if err != nil {
		return nil, err
	}

	operationRecord := domain.Record{
		OperationID:       operation.OperationName,
		UserID:            userID,
		Amount:            cost,
		UserBalance:       user.Balance,
		OperationResponse: result,
	}
	err = uc.operationRepo.RecordOperation(operationRecord)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
