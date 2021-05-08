package domain

//Account represents an account
type Account struct {
	ID                   int64   `json:"account_id"`
	DocumentNumber       string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}
