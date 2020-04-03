package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bienvenidos  a index"))
}

func main() {

	r := chi.NewRouter()
	r.Get("/", index)
	http.ListenAndServe(":3000", r)
}
