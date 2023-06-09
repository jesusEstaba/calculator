package usecase

import (
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/jesusEstaba/calculator/pkg/domain/operations"
)

type CalculatorUseCase struct {
	operationRepo domain.OperationRepository
	userRepo      domain.UserRepository
	operations    map[string]domain.CalculableOperation
}

func NewCalculatorUseCase(
	userRepo domain.UserRepository,
	operationRepo domain.OperationRepository,
	randomStringRepo domain.RandomStringRepository,
) *CalculatorUseCase {
	operationList := make(map[string]domain.CalculableOperation)
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

func (uc *CalculatorUseCase) Calculate(userID string, calculation *domain.Calculation) (*domain.CalculationResult, error) {
	operation, err := uc.operationRepo.GetOperation(calculation.OperationName)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}

	err = user.Withdraw(operation.Cost)
	if err != nil {
		return nil, err
	}

	operationFunc := uc.operations[calculation.OperationName]
	result, err := operationFunc.Calculate(calculation)
	if err != nil {
		return nil, err
	}

	operationRecord := domain.Record{
		OperationID:       operation.ID.Hex(),
		UserID:            userID,
		Amount:            operation.Cost,
		UserBalance:       user.Balance,
		OperationResponse: result,
	}
	err = uc.operationRepo.RecordOperation(operationRecord)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.UpdateUser(*user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
