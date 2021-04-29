package domain

type Account struct {
	Id             uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
