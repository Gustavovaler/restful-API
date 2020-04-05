package main

import (
	"Udemy1/database"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code`
	Description  string `json:"description`
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bienvenidos  a index"))
// }

func main() {

	databaseConnection = database.InitDB()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Get("/products", AllProductos)
	r.Post("/products", Createproducto)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)
	http.ListenAndServe(":3000", r)
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "Successfully deleted"})

}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("UPDATE products set product_code=?, description=? where id=?")
	catch(err)

	_, er := query.Exec(product.Product_Code, product.Description, id)
	catch(er)
	defer query.Close()
	responseWithJSON(w, http.StatusOK, map[string]string{"message": "updated successfully"})

}

func Createproducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)
	query, err := databaseConnection.Prepare("Insert into products (product_code, description) values (?,?)")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()

	responseWithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'') from products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	responseWithJSON(w, 200, products)
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
