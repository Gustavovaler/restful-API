package main

import (
	"Udemy1/database"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	// "net/http"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bienvenidos  a index"))
// }

func main() {

	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	// r := chi.NewRouter()
	// r.Get("/", index)
	// http.ListenAndServe(":3000", r)
}
