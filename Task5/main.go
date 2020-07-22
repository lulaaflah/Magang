package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Root struct {
	XMLName   xml.Name `xml:"Root"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Customers struct {
		Text     string `xml:",chardata"`
		Customer []struct {
			Text         string `xml:",chardata"`
			CustomerID   string `xml:"CustomerID,attr"`
			CompanyName  string `xml:"CompanyName"`
			ContactName  string `xml:"ContactName"`
			ContactTitle string `xml:"ContactTitle"`
			Phone        string `xml:"Phone"`
			FullAddress  struct {
				Text       string `xml:",chardata"`
				Address    string `xml:"Address"`
				City       string `xml:"City"`
				Region     string `xml:"Region"`
				PostalCode string `xml:"PostalCode"`
				Country    string `xml:"Country"`
			} `xml:"FullAddress"`
			Fax string `xml:"Fax"`
		} `xml:"Customer"`
	} `xml:"Customers"`
	Orders struct {
		Text  string `xml:",chardata"`
		Order []struct {
			Text         string `xml:",chardata"`
			CustomerID   string `xml:"CustomerID"`
			EmployeeID   string `xml:"EmployeeID"`
			OrderDate    string `xml:"OrderDate"`
			RequiredDate string `xml:"RequiredDate"`
			ShipInfo     struct {
				Text           string `xml:",chardata"`
				ShippedDate    string `xml:"ShippedDate,attr"`
				ShipVia        string `xml:"ShipVia"`
				Freight        string `xml:"Freight"`
				ShipName       string `xml:"ShipName"`
				ShipAddress    string `xml:"ShipAddress"`
				ShipCity       string `xml:"ShipCity"`
				ShipRegion     string `xml:"ShipRegion"`
				ShipPostalCode string `xml:"ShipPostalCode"`
				ShipCountry    string `xml:"ShipCountry"`
			} `xml:"ShipInfo"`
		} `xml:"Order"`
	} `xml:"Orders"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Root

	if err = xml.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	//Tugas insert kan ke table Customer
	for i := 0; i < len(request.Customers.Customer); i++ {
		customerID := request.Customers.Customer[i].CustomerID
		companyName := request.Customers.Customer[i].CompanyName
		contactName := request.Customers.Customer[i].ContactName
		contactTitle := request.Customers.Customer[i].ContactTitle
		addres := request.Customers.Customer[i].FullAddress.Address
		city := request.Customers.Customer[i].FullAddress.City
		region := request.Customers.Customer[i].FullAddress.Region
		postalCode := request.Customers.Customer[i].FullAddress.PostalCode
		country := request.Customers.Customer[i].FullAddress.Country
		phone := request.Customers.Customer[i].Phone
		fax := request.Customers.Customer[i].Fax

		stmt, err := db.Prepare("INSERT INTO customers (CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
		_, err = stmt.Exec(customerID, companyName, contactName, contactTitle, addres, city, region, postalCode, country, phone, fax)

		if err != nil {
			fmt.Fprintln(w, "Data", customerID, " Duplicate")
		} else {
			fmt.Fprintln(w, "Data", customerID, " Created")
		}

	}

	//Tugas insert kan ke table Order
	for i := 0; i < len(request.Orders.Order); i++ {
		OcustomerID := request.Orders.Order[i].CustomerID
		OemployeeID := request.Orders.Order[i].EmployeeID
		orderDate := request.Orders.Order[i].OrderDate
		requiredDate := request.Orders.Order[i].RequiredDate
		shippedDate := request.Orders.Order[i].ShipInfo.ShippedDate
		shipVia := request.Orders.Order[i].ShipInfo.ShipVia
		freight := request.Orders.Order[i].ShipInfo.Freight
		shipName := request.Orders.Order[i].ShipInfo.ShipName
		shipAddress := request.Orders.Order[i].ShipInfo.ShipAddress
		shipCity := request.Orders.Order[i].ShipInfo.ShipCity
		shipRegion := request.Orders.Order[i].ShipInfo.ShipRegion
		shipPostalCode := request.Orders.Order[i].ShipInfo.ShipPostalCode
		shipCountry := request.Orders.Order[i].ShipInfo.ShipCountry

		stmt, err := db.Prepare("INSERT INTO orders (CustomerID,EmployeeID,OrderDate,RequiredDate,ShippedDate,ShipVia,Freight,ShipName,ShipAddress,ShipCity,ShipRegion,ShipPostalCode,ShipCountry) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
		_, err = stmt.Exec(OcustomerID, OemployeeID, orderDate, requiredDate, shippedDate, shipVia, freight, shipName, shipAddress, shipCity, shipRegion, shipPostalCode, shipCountry)

		if err != nil {
			fmt.Fprintln(w, "Data", OcustomerID, " and Employee ID: ", OemployeeID, " Duplicate")
		} else {
			fmt.Fprintln(w, "Data Customer ID:", OcustomerID, "and Employee ID:", OemployeeID, " Created")
		}
	}
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/customers", getCustomers).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
