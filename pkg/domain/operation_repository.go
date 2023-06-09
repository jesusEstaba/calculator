package domain

type OperationRepository interface {
	GetOperation(operation string) (*Operation, error)
	RecordOperation(Record) error
	GetRecordsByUserAndSearchTermPaginated(search RecordSearch) ([]*Record, error)
}
