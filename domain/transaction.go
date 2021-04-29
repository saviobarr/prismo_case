package domain

type Transaction struct {
	Id            uint64         `json:"transaction_id"`
	OperationType Operation_Type `json:"operation_type"`
	Account       Account        `json:"account"`
	Amount        float64        `json:"amount"`
}
