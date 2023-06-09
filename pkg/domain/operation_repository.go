package domain

type OperationRepository interface {
	GetOperation(operation string) (*Operation, error)
	RecordOperation(Record) error
}
