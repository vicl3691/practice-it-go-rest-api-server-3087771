package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

////HANDLERS
func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is GET\n")
}
func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is POST\n")
}
func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is DELETE\n")
}

func dbSelect() {
	db, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id,name,inventory,price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product:", p.id, " ", p.name, " ", p.inventory, " ", p.price)
	}
}

func MyRouter() {
	dbSelect()

	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server started and listenng at localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
