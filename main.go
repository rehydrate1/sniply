package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Главная"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки"))
}

func createSnippet(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Создание заметки"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("POST /snippet/create", createSnippet)

	port := ":8080"
	log.Println("Запуск сервера на http://localhost:8080")
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}