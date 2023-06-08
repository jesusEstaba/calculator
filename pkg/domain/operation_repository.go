package domain

type CalculatorRepository interface {
	GetOperationCost(operation string) (float64, error)
	RecordOperation(Record) error
}
