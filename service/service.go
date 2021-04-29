package service

import (
	"github.com/saviobarr/prismo_case/dao"
	"github.com/saviobarr/prismo_case/domain"
	"github.com/saviobarr/prismo_case/utils"
)

func CreateTransaction(transaction domain.Transaction) *utils.ApplicationError {

	//compras e saques sempre negativos
	if (transaction.OperationType.Id == 1 || transaction.OperationType.Id == 2 || transaction.OperationType.Id == 3) && (transaction.Amount > 0) {
		transaction.Amount = (transaction.Amount * -1)
	}

	//pagamentos sempre positivos
	if transaction.OperationType.Id == 4 && transaction.Amount < 0 {
		transaction.Amount = (transaction.Amount * -1)
	}
	return dao.CreateTransaction(transaction)

}

func CreateAccount(account domain.Account) *utils.ApplicationError {
	return dao.CreateAccount(account)

}

func GetAccount(id int) (*domain.Account, *utils.ApplicationError) {
	return dao.GetAccount(id)

}
