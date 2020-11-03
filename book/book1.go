package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// var books []Book

// type Book struct {
// 	Id     string  `json: "id"`
// 	Name   string  `json: "name"`
// 	BookNo string  `json: "bookno"`
// 	Author *Author `json: "author"`
// }

// type Author struct {
// 	Firstname string `json: "firstname"`
// 	Lastname  string `json: "lastname"`
// }

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// func main() {
// 	r := mux.NewRouter()

// 	books = append(books, Book{Id: "1", Name: "Book one", BookNo: "123", Author: &Author{Firstname: "Triyank", Lastname: "Jain"}})
// 	fmt.Println(books)
// 	r.HandleFunc("/api/books", getBooks).Methods("GET")

// }
