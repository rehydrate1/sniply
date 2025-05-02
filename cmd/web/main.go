package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/{id}", showSnippet)
	mux.HandleFunc("POST /snippet/create", createSnippet)

	port := ":8080"
	log.Println("Запуск сервера на http://localhost:8080")
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}