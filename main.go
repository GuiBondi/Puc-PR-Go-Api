package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Response struct {
	Message string `json: "message"`
}


func  greetings(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Greetings!!"}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}


func main(){
	http.HandleFunc("/hello", greetings)

	log.Println("Iniciando server na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

