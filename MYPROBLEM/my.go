package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

//fjbhv
type Order struct {
	OrderID      int    `json: "orderid" `
	CustomerName string `json: "customername"`
}

func main() {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/triyank?charset=utf8&parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect Database")
	} else {
		fmt.Println("Database stabilised")
	}

	handleRequests()
	defer db.Close()

}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:3000/")
	log.Println("Quit the server with CONTROL-C.")

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getOrders).Methods("GET")
	router.HandleFunc("/all", returnAllBookings).Methods("GET")
	router.HandleFunc("/new", createOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", router))
}
func getOrders(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	welcome := "hello world"
	json.NewEncoder(res).Encode(welcome)
}

func createOrder(res http.ResponseWriter, req *http.Request) {

	order := Order{}
	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &order)
	db.Create(&order)
	json.NewEncoder(res).Encode(order)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookings := []Order{}
	//retrieve all bookings from db
	db.Find(&bookings)
	json.NewEncoder(w).Encode(bookings)
}
