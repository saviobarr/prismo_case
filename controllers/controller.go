package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saviobarr/prismo_case/domain"
	"github.com/saviobarr/prismo_case/service"
	"github.com/saviobarr/prismo_case/utils"
)

func CreateTransaction(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")
	var transaction domain.Transaction
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&transaction)

	if err != nil {
		http.Error(resp, "Bad Request", http.StatusBadRequest)
		return
	}

	apiErr := service.CreateTransaction(transaction)

	if apiErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(resp)
		encoder.Encode(apiErr)
	}

	resp.WriteHeader(http.StatusOK)

	jsonValue, _ := json.Marshal(utils.AppMsgs{http.StatusOK, "Transaction was recorded"})

	resp.Write(jsonValue)

}

func CreateAccount(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")

	var account domain.Account
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&account)

	if err != nil {
		http.Error(resp, "Bad Request", http.StatusBadRequest)
		return
	}

	apiErr := service.CreateAccount(account)

	if apiErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(resp)
		encoder.Encode(apiErr)
	}

	resp.WriteHeader(http.StatusOK)

	jsonValue, _ := json.Marshal(utils.AppMsgs{http.StatusOK, "Account was created"})

	resp.Write(jsonValue)
}

func GetAccount(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	resp.Header().Add("Content-Type", "application/json")

	//test if param is a number
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "account_id must be a number/cannot be empty",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)

		return
	}

	account, apiErr := service.GetAccount(id)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)

		return
	}

	resp.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(resp)
	encoder.Encode(account)

}
