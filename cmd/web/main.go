package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	books = append(books, Book{Id: 1, Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{Id: 2, Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	muxServer := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	muxServer.Handle("/static/", http.StripPrefix("/static", fileServer))

	//r := mux.Router{}
	//muxServer.HandleFunc("/books/{id}", getBook)

	muxServer.HandleFunc("/books/all", getBooks)
	muxServer.HandleFunc("/books/book", getBook)
	muxServer.HandleFunc("/books/create", createBook)
	muxServer.HandleFunc("/books/delete", deleteBook)
	muxServer.HandleFunc("/books/update", updateBook)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  muxServer,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
