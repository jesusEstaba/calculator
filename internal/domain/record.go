package domain

type Record struct {
	ID                string
	OperationID       string
	UserID            string
	Amount            float64
	UserBalance       float64
	OperationResponse any
	Date              string
}
