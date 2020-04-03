package database

import (
	"database/sql"
)

func InitDB() *sql.DB {

	connection := "root:@/northwind"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	return db
}
