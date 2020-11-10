package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}
type Order struct {
	OrderID      int    `json: "orderid" `
	CustomerName string `json: "customername"`
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:3000/")
	log.Println("Quit the server with CONTROL-C.")

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	// myRouter.HandleFunc("/booking/{id}", returnSingleBooking).Methods("GET")
	// myRouter.HandleFunc("/all", returnAllBookings).Methods("GET")
	// myRouter.HandleFunc("/update/{id}", updateBooking).Methods("PUT")
	// myRouter.HandleFunc("/delete/{id}", deleteBooking).Methods("DELETE")
	// myRouter.HandleFunc("/new", createNewBooking).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func main() {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/triyank?charset=utf8&parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Order{})
	handleRequests()
	defer db.Close()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Booking Home Page!")
	fmt.Println("Endpoint Hit: Booking Home Page")
}

// func createNewBooking(w http.ResponseWriter, r *http.Request) {

// 	return the string response containing the request body
// 		reqBody, _ := ioutil.ReadAll(r.Body)
// 		var booking Booking
// 		json.Unmarshal(reqBody, &booking)
// 		db.Create(&booking)
// 		json.NewEncoder(w).Encode(booking)
// 	}
// 	var order Order
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	json.Unmarshal(reqBody, &order)
// 	db.Create(&order)
// 	json.NewEncoder(w).Encode(order)
// }

// func returnAllBookings(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	bookings := []Order{}
// 	//retrieve all bookings from db
// 	db.Find(&bookings)
// 	json.NewEncoder(w).Encode(bookings)
// }

// func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	bookings := []Booking{}
// 	db.Find(&bookings)
// 	for _, booking := range bookings {
// 		//string to int
// 		s, err := strconv.Atoi(key)
// 		if err == nil {
// 			if booking.Id == s {
// 				fmt.Println(booking)
// 				json.NewEncoder(w).Encode(booking)
// 			}
// 		}
// 	}
// }

// func updateBooking(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	bookings := []Booking{}
// 	db.Find(&bookings)
// 	s, err := strconv.Atoi(key)
// 	if s > len(bookings) {
// 		fmt.Fprintln(w, "Id doesn't exist")
// 		return
// 	}
// 	for _, booking := range bookings {
// 		//string to int

// 		if err == nil {
// 			if booking.Id == s {
// 				reqBody, _ := ioutil.ReadAll(r.Body)
// 				json.Unmarshal(reqBody, &booking)
// 				fmt.Println(booking)
// 				db.Save(booking)
// 				json.NewEncoder(w).Encode(booking)
// 			}
// 		}
// 	}
// }

// func deleteBooking(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	bookings := []Booking{}
// 	db.Find(&bookings)
// 	for _, booking := range bookings {
// 		//string to int
// 		s, err := strconv.Atoi(key)
// 		if err == nil {
// 			if booking.Id == s {
// 				db.Table("bookings").Where("id= ?", s).Delete(&Booking{})
// 				json.NewEncoder(w).Encode(booking)
// 			}
// 		}
// 	}
// }
