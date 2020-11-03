package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("App is Served")

	router.HandleFunc("/coffee", getList).Methods("GET")
	router.HandleFunc("/coffee", addShop).Methods("POST")
	router.HandleFunc("/coffee/{id}", updateShop).Methods("PUT")
	router.HandleFunc("/coffee/{id}", getShop).Methods("GET")
	router.HandleFunc("/coffee/{id}", deleteShop).Methods("DELETE")

	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(":1200", router))
}
