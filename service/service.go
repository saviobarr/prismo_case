package service

import (
	"github.com/saviobarr/prismo_case/dao"
	"github.com/saviobarr/prismo_case/domain"
	"github.com/saviobarr/prismo_case/utils"
)

func CreateTransaction(transaction domain.Transaction) *utils.ApplicationError {
	return dao.CreateTransaction(transaction)

}

func CreateAccount(account domain.Account) *utils.ApplicationError {
	return dao.CreateAccount(account)

}

func GetAccount(id int) (*domain.Account, *utils.ApplicationError) {
	return dao.GetAccount(id)

}
