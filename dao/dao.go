package dao

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/saviobarr/prismo_case/domain"
	"github.com/saviobarr/prismo_case/utils"
)

//CreateTransaction creates a transaction into database
func CreateTransaction(t domain.Transaction) *utils.ApplicationError {
	query := "INSERT INTO TRANSACTION (ACCOUNT_ID,OPERATION_TYPE_ID,AMOUNT,EVENT_DATE) VALUES(?,?,?,?)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error trying to connect to DB %v: ", err.Error()),
			StatusCode: -1,
			Code:       "error",
		}

	}

	defer stmt.Close()
	_, ex := stmt.Exec(t.Account.ID, t.OperationType.Id, t.Amount, time.Now().Format("2006-01-02 15:04:05"))

	if ex != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error creating record in DB %v: ", ex.Error()),
			StatusCode: -1,
			Code:       "error",
		}
	}

	return nil

}

//CreateAccount creates an account into database
func CreateAccount(account domain.Account) *utils.ApplicationError {
	query := "INSERT INTO ACCOUNTS (document_number,available_credit_limit) values (?,?)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error trying to connect to DB %v: ", err.Error()),
			StatusCode: -1,
			Code:       "error",
		}

	}

	defer stmt.Close()
	_, ex := stmt.Exec(account.DocumentNumber, account.AvailableCreditLimit)

	if ex != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error creating record in DB %v: ", ex.Error()),
			StatusCode: -1,
			Code:       "error",
		}
	}

	return nil

}

//GetAccount gets one account, given an id
func GetAccount(id int64) (*domain.Account, *utils.ApplicationError) {

	query := "SELECT * FROM ACCOUNTS WHERE account_id = ?"

	db := get()
	defer db.Close()

	row := db.QueryRow(query, id)
	var account domain.Account
	switch err := row.Scan(&account.ID, &account.DocumentNumber, &account.AvailableCreditLimit); err {
	case sql.ErrNoRows:
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Account %v was not found", id),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	case nil:
		return &account, nil
	default:
		panic(err)

	}

}

//UpdateAccountTransaction creates an account into database
func UpdateAccountTransaction(account *domain.Account, transaction domain.Transaction) *utils.ApplicationError {
	queryInsertTransaction := "INSERT INTO TRANSACTION (ACCOUNT_ID,OPERATION_TYPE_ID,AMOUNT,EVENT_DATE) VALUES(?,?,?,?)"
	queryUpdateAccount := "UPDATE ACCOUNTS SET available_credit_limit = ? WHERE account_id = ?"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(queryUpdateAccount)

	if err != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error trying to connect to DB %v: ", err.Error()),
			StatusCode: -1,
			Code:       "error",
		}

	}

	_, ex := stmt.Exec(account.AvailableCreditLimit, account.ID)

	if ex != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error creating record in DB %v: ", ex.Error()),
			StatusCode: -1,
			Code:       "error",
		}
	}

	//inserir transacao
	stmt1, err1 := db.Prepare(queryInsertTransaction)
	if err1 != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error trying to connect to DB %v: ", err1.Error()),
			StatusCode: -1,
			Code:       "error",
		}

	}

	_, ex1 := stmt1.Exec(transaction.Account.ID, transaction.OperationType.Id, transaction.Amount, time.Now().Format("2006-01-02 15:04:05"))

	if ex1 != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error trying to save in DB %v: ", ex1.Error()),
			StatusCode: -1,
			Code:       "error",
		}

	}

	defer stmt.Close()
	defer stmt1.Close()
	return nil

}
