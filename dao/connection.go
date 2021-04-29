package dao

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/saviobarr/prismo_case/utils"
)

func get() *sql.DB {
	config, err := utils.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
