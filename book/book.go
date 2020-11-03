package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Book

type Book struct {
	Id     string  `json: "id"`
	BookNo string  `json: "bookno"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname  string `json: "firstname"`
	Secondname string `json: "secondname"`
}

//to get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//to get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Book{})
}

//to create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//to  update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	params := mux.Vars(r)
	for i, v := range books {
		if v.Id == params["id"] {
			book.Id = params["id"]
			books[i] = book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//to delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, val := range books {
		if val.Id == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main() {

	r := mux.NewRouter()

	books = append(books, Book{Id: "1", BookNo: "123", Title: "The curse of The Black Pearl", Author: &Author{Firstname: "Jack", Secondname: "Sparrow"}})
	books = append(books, Book{Id: "2", BookNo: "234", Title: "Dead Man's Chest", Author: &Author{Firstname: "Hector", Secondname: "Barbossa"}})
	books = append(books, Book{Id: "3", BookNo: "345", Title: "At World's End", Author: &Author{Firstname: "Elizabeth", Secondname: "Swann"}})
	books = append(books, Book{Id: "4", BookNo: "456", Title: "On Stranger's Tides", Author: &Author{Firstname: "William", Secondname: "Turner"}})
	fmt.Println(books)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1800", r))
}
