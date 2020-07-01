package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type Product struct {
	ProductID       string `json:"ProductID"`
	ProductName     string `json:"ProductName"`
	SupplierID      string `json:"SupplierID"`
	CategoryID      string `json:"CategoryID"`
	QuantityPerUnit string `json:"QuantityPerUnit"`
	UnitPrice       string `json:"UnitPrice"`
	UnitsInStock    string `json:"UnitsInStock"`
	UnitsOnOrder    string `json:"UnitsOnOrder"`
	ReorderLevel    string `json:"ReorderLevel"`
	Discontinued    string `json:"Discontinued"`
	Description     string `json:"Description"`
}

// Get all orders

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []Product

	sql := `SELECT
				ProductID,
				IFNULL(ProductName,''),
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitsInStock,
				IFNULL(UnitsOnOrder,'') UnitsOnOrder,
				IFNULL(ReorderLevel,'') ReorderLevel,
				IFNULL(Discontinued,'') Discontinued,
				IFNULL(Description,'') Description
			FROM products`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var product Product
		err := result.Scan(&product.ProductID, &product.ProductName, &product.SupplierID, &product.CategoryID, &product.QuantityPerUnit, &product.UnitPrice, &product.UnitsInStock, &product.UnitsOnOrder, &product.ReorderLevel, &product.Discontinued, &product.Description)

		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		ProductID := r.FormValue("productID")
		ProductName := r.FormValue("productName")
		SupplierID := r.FormValue("supplierID")
		CategoryID := r.FormValue("categoryID")
		QuantityPerUnit := r.FormValue("quantityPerUnit")
		UnitPrice := r.FormValue("unitPrice")
		UnitsInStock := r.FormValue("unitsInStock")
		UnitsOnOrder := r.FormValue("unitsOnOrder")
		ReorderLevel := r.FormValue("reorderLevel")
		Discontinued := r.FormValue("discontinued")
		Description := r.FormValue("description")

		stmt, err := db.Prepare("INSERT INTO products (ProductID, ProductName, SupplierID, CategoryID, QuantityPerUnit, UnitPrice, UnitsInStock, UnitsOnOrder, ReorderLevel, Discontinued, Description) VALUES(?,?,?,?,?,?,?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(ProductID, ProductName, SupplierID, CategoryID, QuantityPerUnit, UnitPrice, UnitsInStock, UnitsOnOrder, ReorderLevel, Discontinued, Description)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {

			fmt.Fprintf(w, "Date Created")
			//http.Redirect(w, r, "/", 301)
		}

	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	params := mux.Vars(r)

	sql := `SELECT
				ProductID,
				IFNULL(ProductName,''),
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitsInStock,
				IFNULL(UnitsOnOrder,'') UnitsOnOrder,
				IFNULL(ReorderLevel,'') ReorderLevel,
				IFNULL(Discontinued,'') Discontinued,
				IFNULL(Description,'') Description
			FROM products WHERE ProductID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var product Product

	for result.Next() {

		err := result.Scan(&product.ProductID, &product.ProductName, &product.SupplierID, &product.CategoryID, &product.QuantityPerUnit, &product.UnitPrice, &product.UnitsInStock, &product.UnitsOnOrder, &product.ReorderLevel, &product.Discontinued, &product.Description)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newProductName := r.FormValue("ProductName")

		stmt, err := db.Prepare("UPDATE products SET ProductName = ? WHERE ProductID = ?")

		_, err = stmt.Exec(newProductName, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Product with ProductID = %s was updated", params["id"])
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM products WHERE ProductID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Product with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var products []Product

	ProductID := r.FormValue("productID")
	ProductName := r.FormValue("productName")

	sql := `SELECT
				ProductID,
				IFNULL(ProductName,''),
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitsInStock,
				IFNULL(UnitsOnOrder,'') UnitsOnOrder,
				IFNULL(ReorderLevel,'') ReorderLevel,
				IFNULL(Discontinued,'') Discontinued,
				IFNULL(Description,'') Description
			FROM products WHERE ProductID = ? AND ProductName = ?`

	result, err := db.Query(sql, ProductID, ProductName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var product Product

	for result.Next() {

		err := result.Scan(&product.ProductID, &product.ProductName, &product.SupplierID, &product.CategoryID, &product.QuantityPerUnit, &product.UnitPrice, &product.UnitsInStock, &product.UnitsOnOrder, &product.ReorderLevel, &product.Discontinued, &product.Description)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	//new
	r.HandleFunc("/getproduct", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8282", r))
}
