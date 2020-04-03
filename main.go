package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/go-chi/chi"
	// "net/http"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bienvenidos  a index"))
// }

func main() {

	connection := "root:@/northwind"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// r := chi.NewRouter()
	// r.Get("/", index)
	// http.ListenAndServe(":3000", r)
}
