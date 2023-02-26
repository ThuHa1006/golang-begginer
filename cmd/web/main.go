package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	muxServer := http.NewServeMux()
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//r := mux.Router{}
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	muxServer.Handle("/static/", http.StripPrefix("/static", fileServer))

	books = append(books, Book{Id: 1, Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{Id: 2, Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	muxServer.HandleFunc("/books/all", getBooks)
	muxServer.HandleFunc("/books/book", getBook)
	//muxServer.HandleFunc("/books/{id}", getBook)
	muxServer.HandleFunc("/books/create", createBook)
	muxServer.HandleFunc("/books/delete", deleteBook)
	muxServer.HandleFunc("/books/update", updateBook)
	log.Print("Starting server on ", *addr)
	err := http.ListenAndServe(*addr, muxServer)
	log.Fatal(err)
}
