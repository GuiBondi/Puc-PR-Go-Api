package main

import (
    "encoding/json"
    "log"
    "net/http"
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

// type Response struct {
// 	Message string `json: "message"`
// }


// func  greetings(w http.ResponseWriter, r *http.Request) {
// 	response := Response{Message: "Greetings!!"}

// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(response)
// }


// func main(){
// 	http.HandleFunc("/hello", greetings)

// 	log.Println("Iniciando server na porta :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

