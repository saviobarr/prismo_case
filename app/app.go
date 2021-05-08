package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saviobarr/prismo_case/controllers"
)

//StartApp Start listener and routers
func StartApp() {

	r := mux.NewRouter()
	r.HandleFunc("/transaction", controllers.CreateTransaction).Methods("POST")
	r.HandleFunc("/account", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("/account/{id}", controllers.GetAccount).Methods("GET")
	r.HandleFunc("/account", controllers.GetAccount).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
