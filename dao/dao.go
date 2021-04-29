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

//Creates a transaction into database
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
	_, ex := stmt.Exec(t.Account.Id, t.OperationType.Id, t.Amount, time.Now().Format("2006-01-02 15:04:05"))

	if ex != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error creating record in DB %v: ", ex.Error()),
			StatusCode: -1,
			Code:       "error",
		}
	}

	return nil

}

//Creates an account into database
func CreateAccount(account domain.Account) *utils.ApplicationError {
	query := "INSERT INTO ACCOUNTS (document_number) values (?)"
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
	_, ex := stmt.Exec(account.DocumentNumber)

	if ex != nil {
		return &utils.ApplicationError{
			Message:    fmt.Sprintf("Error creating record in DB %v: ", ex.Error()),
			StatusCode: -1,
			Code:       "error",
		}
	}

	return nil

}

//gets one account, given an id
func GetAccount(id int) (*domain.Account, *utils.ApplicationError) {
	query := "SELECT * FROM ACCOUNTS WHERE account_id = ?"

	db := get()
	defer db.Close()

	row := db.QueryRow(query, id)
	var account domain.Account
	switch err := row.Scan(&account.Id, &account.DocumentNumber); err {
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
