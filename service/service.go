package service

import (
	"fmt"
	"net/http"

	"github.com/saviobarr/prismo_case/dao"
	"github.com/saviobarr/prismo_case/domain"
	"github.com/saviobarr/prismo_case/utils"
)

//CreateTransaction performs business rules implementation and propagates to DAO layer
func CreateTransaction(transaction domain.Transaction) *utils.ApplicationError {

	var account, err = dao.GetAccount(transaction.Account.ID)
	if err != nil {

		return &utils.ApplicationError{
			Message:    fmt.Sprintf("An error ocurrer trying to retrieve available limit"),
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_server_error",
		}

	}

	//compras e saques sempre negativos
	if (transaction.OperationType.Id == 1 || transaction.OperationType.Id == 2 || transaction.OperationType.Id == 3) && (transaction.Amount > 0) {
		transaction.Amount = (transaction.Amount * -1)
	}

	//pagamentos sempre positivos
	if transaction.OperationType.Id == 4 && transaction.Amount < 0 {
		transaction.Amount = (transaction.Amount * -1)
	}

	//verificar limite disponivel
	tamount := account.AvailableCreditLimit + transaction.Amount
	if tamount < 0 {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Sorry, this transaction amount is over your credit limit "),
			StatusCode: http.StatusAccepted,
			Code:       "NOK",
		}

	}

	account.AvailableCreditLimit = tamount

	return dao.UpdateAccountTransaction(account, transaction)

}

//CreateAccount performs business rules implementation and propagates to DAO layer
func CreateAccount(account domain.Account) *utils.ApplicationError {
	return dao.CreateAccount(account)

}

//GetAccount performs business rules implementation and propagates to DAO layer
func GetAccount(id int64) (*domain.Account, *utils.ApplicationError) {
	return dao.GetAccount(id)

}
