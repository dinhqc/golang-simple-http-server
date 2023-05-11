package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 color=blue>Hello</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := Book{
		Title:  "Book title",
		Author: "This author",
		Pages:  125,
	}

	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
