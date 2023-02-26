package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Allow", http.MethodGet)
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Allow", http.MethodGet)
	//vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["id"])
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	for _, item := range books {
		if item.Id == id {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			break
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Allow", http.MethodPost)
	var book Book
	var maxId = books[0].Id
	for _, item := range books {
		if item.Id > maxId {
			maxId = item.Id
		}
	}
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = maxId + 1
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Allow", http.MethodDelete)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	for index, item := range books {
		if item.Id == id {
			books = append(books[:index], books[index+1:]...)
			err := json.NewEncoder(w).Encode(books)
			if err != nil {
				return
			}
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Allow", http.MethodPut)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	for index, item := range books {
		if item.Id == id {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = id
			books = append(books, book)
			err := json.NewEncoder(w).Encode(books)
			if err != nil {
				return
			}
			break
		}
	}
}
