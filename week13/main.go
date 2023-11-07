package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "M@dhu12345"
	dbName := "WEEK13"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}
