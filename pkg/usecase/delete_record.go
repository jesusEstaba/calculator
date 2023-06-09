package usecase

import "github.com/jesusEstaba/calculator/pkg/domain"

type DeleteRecordUseCase struct {
	operationRepo domain.OperationRepository
}

func NewDeleteRecordUseCase(
	operationRepo domain.OperationRepository,
) *DeleteRecordUseCase {
	return &DeleteRecordUseCase{
		operationRepo,
	}
}

func (uc *DeleteRecordUseCase) Delete(userID string, id string) error {
	return uc.operationRepo.DeleteFromUser(userID, id)
}
