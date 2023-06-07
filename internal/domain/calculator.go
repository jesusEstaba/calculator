package domain

type Calculation struct {
	OperationName string  `json:"operation"`
	OperandA      float64 `json:"a"`
	OperandB      float64 `json:"b"`
}

type CalculationResult struct {
	Result string `json:"result"`
}

type Operation interface {
	Calculate(*Calculation) (*CalculationResult, error)
}
