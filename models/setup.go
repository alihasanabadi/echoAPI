package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456."
	DB_NAME     = "ginDB"
	DB_HOST     = "localhost"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}
