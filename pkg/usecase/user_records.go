package usecase

import "github.com/jesusEstaba/calculator/pkg/domain"

type SearchUserRecordsUseCase struct {
	operationRepo domain.OperationRepository
}

func NewSearchUserRecordsUseCase(
	operationRepo domain.OperationRepository,
) *SearchUserRecordsUseCase {
	return &SearchUserRecordsUseCase{
		operationRepo,
	}
}

func (uc *SearchUserRecordsUseCase) Search(search domain.RecordSearch) ([]*domain.Record, error) {
	return uc.operationRepo.GetRecordsByUserAndSearchTermPaginated(search)
}
