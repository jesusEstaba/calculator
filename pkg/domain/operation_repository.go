package domain

type OperationRepository interface {
	GetOperation(operation string) (*Operation, error)
	DeleteFromUser(userID string, id string) error
	RecordOperation(Record) error
	GetRecordsByUserAndSearchTermPaginated(search RecordSearch) ([]*Record, error)
}
