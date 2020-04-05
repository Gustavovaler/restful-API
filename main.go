package main

import (
	"Udemy1/database"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	// "net/http"
)

func main() {

	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

}
