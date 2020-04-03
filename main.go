package main

import (
	"Udemy1/database"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/go-chi/chi"
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
