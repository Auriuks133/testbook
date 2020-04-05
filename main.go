package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book information
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	IsBn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

// Author iformation
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	return
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range books {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//MaxID return the biggest id in the books array
func MaxID() string {
	var max int64 = 0
	for _, book := range books {
		parsedID, _ := strconv.ParseInt(book.ID, 0, 8)
		if max < parsedID {
			max = parsedID
		}
	}
	max++
	return strconv.FormatInt(max, 10)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = MaxID()
	books = append(books, newBook)
	return
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	i := findBook(id)

	if i == -1 {
		return
	}

	removeFromBooks(id)
	books = append(books, newBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	removeFromBooks(id)
}

func findBook(index string) int {
	i := -1
	for i2, book := range books {
		if index == book.ID {
			i = i2
		}
	}

	return i
}

func removeFromBooks(index string) {
	i := findBook(index)

	if i == -1 {
		return
	}

	books[i] = books[len(books)-1]
	books[len(books)-1] = *new(Book)
	books = books[:len(books)-1]
}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Tile1", IsBn: "5555", Author: &Author{FirstName: "Jon", LastName: "Smith"}})
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8888", r))
}
