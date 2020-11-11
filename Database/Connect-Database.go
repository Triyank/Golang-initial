package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	laptop []Laptop
	db     *gorm.DB
	err    error
)

type Laptop struct {
	ID      int    `json: "id"; gorm:"primary_key; AUTO_INCREMENT"`
	Brand   string `json: "brand"`
	Windows string `json: "windows"`
}

func main() {

	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/triyank?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}
	fmt.Println("SQL tutorial")

	db.AutoMigrate(&Laptop{})
	handleRequests()

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/l", getLaptops).Methods("GET")
	router.HandleFunc("/l/{id}", getLaptop).Methods("GET")
	router.HandleFunc("/l", addLaptop).Methods("POST")
	router.HandleFunc("/l/{id}", deleteLaptop).Methods("DELETE")
	router.HandleFunc("/l/{id}", updateLaptop).Methods("PUT")

	log.Fatal(http.ListenAndServe(":1500", router))
}

func homePage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	welcome := "Welcome to the Home Page."
	json.NewEncoder(res).Encode(welcome)
}

func getLaptops(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	db.Find(&laptop)
	json.NewEncoder(res).Encode(laptop)
}

func getLaptop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	count := 0
	params := mux.Vars(req)
	s, _ := strconv.Atoi(params["id"])
	for _, v := range laptop {
		if v.ID == s {
			json.NewEncoder(res).Encode(v)
			count = count + 1
			return
		}
	}
	if count == 0 {
		fail := "Incorrect ID"
		json.NewEncoder(res).Encode(fail)
	}
}

func addLaptop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var lappy Laptop
	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &lappy)

	db.Create(&lappy)

	json.NewEncoder(res).Encode(lappy)
}

func deleteLaptop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	count := 0
	s, _ := strconv.Atoi(params["id"])
	for i, v := range laptop {
		if v.ID == s {
			laptop = append(laptop[:i], laptop[i+1:]...)
			count = count + 1
			json.NewEncoder(res).Encode(laptop)
			return
		}
	}
	if count == 0 {
		fail := "Incorrect ID"
		json.NewEncoder(res).Encode(fail)
	}
}

func updateLaptop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	count := 0
	s, _ := strconv.Atoi(params["id"])
	for i, v := range laptop {
		if v.ID == s {
			var newLappy Laptop
			json.NewDecoder(req.Body).Decode(&newLappy)
			newLappy.ID = v.ID
			laptop[i] = newLappy
			json.NewEncoder(res).Encode(laptop[i])
			count = count + 1
			return
		}
	}
	if count == 0 {
		fail := "Incorrect ID"
		json.NewEncoder(res).Encode(fail)
	}
}
