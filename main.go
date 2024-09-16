package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
var books = []Book{
	{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: 1951},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
	{ID: 3, Title: "1984", Author: "George Orwell", Year: 1949},
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ID from the URL
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Search for the book with the given ID
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// If no book is found
	http.Error(w, "Book not found", http.StatusNotFound)
}
func main() {
	// Route to get all books
	http.HandleFunc("/books", getBooksHandler)

	// Route to get a book by ID (e.g., /books/1)
	http.HandleFunc("/books/", getBookByIDHandler)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

