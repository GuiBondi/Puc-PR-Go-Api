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
	{ID: 1, Title: "O Apanhador no Campo de Centeio", Author: "J.D. Salinger", Year: 1951},
	{ID: 2, Title: "O Sol é Para Todos", Author: "Harper Lee", Year: 1960},
	{ID: 3, Title: "1984", Author: "George Orwell", Year: 1949},
}


func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	idStr := r.URL.Path[len("/livros/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id do livro invalido", http.StatusBadRequest)
		return
	}

	
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	
	http.Error(w, "Livro nao encontrado", http.StatusNotFound)
}
func main() {
	
	http.HandleFunc("/livros", getBooksHandler)

	
	http.HandleFunc("/livros/", getBookByIDHandler)

	
	fmt.Println("Server rodando na porta http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

