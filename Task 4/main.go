package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

//struct modif
type Modification struct {
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	GUID    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	ID          int64  `json:"id"`
	Link        string `json:"link"`
	Modified    string `json:"modified"`
	ModifiedGmt string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Type string `json:"type"`
}

func inputModif(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Modification

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Tanggal := request.Date
	TglGMT := request.DateGmt
	RenderGUID := request.GUID.Rendered
	Id := request.ID
	link := request.Link
	Modifikasi := request.Modified
	ModifGMT := request.ModifiedGmt
	slug := request.Slug
	RenderTitle := request.Title.Rendered
	Tipe := request.Type

	//Insert to Database
	stmt, err := db.Prepare("INSERT INTO modification (Date,DateGmt, RenderedGUID, ID, Link, Modified, ModifiedGmt, Slug, RenderedTitle, Type) VALUES(?,?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(Tanggal, TglGMT, RenderGUID, Id, link, Modifikasi, ModifGMT, slug, RenderTitle, Tipe)
	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

//struct thumbnail

type Thumbnail struct {
	Colors struct {
		Category string `json:"category"`
		Code     struct {
			Hex  string `json:"hex"`
			Rgba string `json:"rgba"`
		} `json:"code"`
		Color string `json:"color"`
		Type  string `json:"type"`
	} `json:"colors"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}

func inputThumbnail(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Thumbnail

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Kategori := request.Colors.Category
	Heks := request.Colors.Code.Hex
	R_GBA := request.Colors.Code.Rgba
	Warna := request.Colors.Color
	Tipe := request.Colors.Type
	Tinggi := request.Thumbnail.Height
	Link := request.Thumbnail.URL
	Lebar := request.Thumbnail.Width

	//Tugas insert kan ke table thumbnail
	stmt, err := db.Prepare("INSERT INTO thumbnail (Color,Category,Hex,Rgba,Type,URL,Width,Height) VALUES(?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(Kategori, Heks, R_GBA, Warna, Tipe, Tinggi, Link, Lebar)
	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tugas4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/inputModif", inputModif).Methods("POST")
	r.HandleFunc("/inputThumbnail", inputThumbnail).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
