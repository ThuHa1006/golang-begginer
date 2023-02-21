package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//// bai 1: random so trong khoang min max
	//fmt.Println("Bai 1: random number in a range")
	//random(1, 50)
	//
	//// bai 2 + 3
	//fmt.Printf("\n")
	//fmt.Println("Bai 2 + 3: countdown from Tet")
	//countdown(time.Date(2023, 1, 27, 0, 0, 0, 0, time.Local))
	//
	//// bai 4
	//fmt.Printf("\n")
	//fmt.Println("Bai 4: Upper the first letter in sting")
	//fmt.Println(upperFirstCharacter("qe TaLeNts"))
	//
	//// bai 5
	//fmt.Printf("\n")
	//fmt.Println("Bai 5")
	//str := "File1.docx.txt"
	//split := strings.Split(str, ".")
	//fmt.Println(split[len(split)-1])
	//
	//// bai 6
	//fmt.Printf("\n")
	//fmt.Println("Bai 6")
	//fmt.Println(convertNumberToArray(5678972))
	//
	//// bai 7
	//fmt.Printf("\n")
	//fmt.Println("Bai 7:")
	//fmt.Println(generatePassword(8, 1, 1, 1, 1))
	//
	//// bai 8
	//fmt.Printf("\n")
	//fmt.Println("Bai 8:")
	//var data []PersonalInfo
	//
	//person1 := PersonalInfo{
	//	id:   1,
	//	name: "Truong",
	//	age:  30,
	//}
	//person2 := PersonalInfo{
	//	id:   2,
	//	name: "Phuc",
	//	age:  40,
	//}
	//data = append(data, person1)
	//data = append(data, person2)
	//exportFile("test.txt", data)
	//
	//dataCSV := [][]string{
	//	{"id", "name", "age"},
	//	{"1", "Truong", "10"},
	//	{"2", "Phuc", "20"},
	//}
	//exportFileCSV("test.csv", dataCSV)

	muxServer := http.NewServeMux()
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
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", muxServer)
	log.Fatal(err)
}

func random(min, max int) {
	fmt.Println(rand.Intn(max-min) + min)
}

func countdown(t time.Time) {
	currentTime := time.Now()
	distance := t.Sub(currentTime)
	total := int(distance.Seconds())

	days := total / (60 * 60 * 24)
	hours := total / (60 * 60) % 24
	minutes := (total / 60) % 60
	seconds := total % 60

	fmt.Println("còn ", days, " ngày ", hours, " giờ ", minutes, " phút ", seconds, " nữa là đến tết")

	time.Sleep(1 * time.Second)
	if total > 0 {
		countdown(t)
	}
}

func upperFirstCharacter(str string) string {
	result := strings.ToLower(str)
	firstLetter := strings.ToUpper(string(result[0]))
	return strings.Replace(result, string(result[0]), firstLetter, 1)
}

func convertNumberToArray(number int) []string {
	str := strconv.Itoa(number)
	var arrString []string

	for i := 0; i < len(str); i++ {
		arrString = append(arrString, string(str[i]))
	}
	return arrString
}

func generatePassword(passwordLength, minLower, minUpperCase, minNumber, minSpecialChar int) string {
	var arrPassword []string
	var (
		lowerChar   = "abcdedfghijklmnopqrst"
		upperChar   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialChar = "~!@#$%^&*()"
		number      = "0123456789"
		all         = lowerChar + upperChar + specialChar + number
	)

	// set upper case
	for i := 0; i < minUpperCase; i++ {
		indexChar := rand.Intn(len(upperChar))
		arrPassword = append(arrPassword, string(upperChar[indexChar]))
	}

	// set lower case
	for i := 0; i < minLower; i++ {
		indexChar := rand.Intn(len(lowerChar))
		arrPassword = append(arrPassword, string(lowerChar[indexChar]))
	}

	// set special char
	for i := 0; i < minSpecialChar; i++ {
		indexChar := rand.Intn(len(specialChar))
		arrPassword = append(arrPassword, string(specialChar[indexChar]))
	}

	// set number
	for i := 0; i < minNumber; i++ {
		indexChar := rand.Intn(len(number))
		arrPassword = append(arrPassword, string(number[indexChar]))
	}

	// set all char
	remaining := passwordLength - minLower - minUpperCase - minNumber - minSpecialChar
	for i := 0; i < remaining; i++ {
		indexChar := rand.Intn(len(all))
		arrPassword = append(arrPassword, string(all[indexChar]))
	}
	return strings.Join(arrPassword, "")
}

type PersonalInfo struct {
	id   int
	name string
	age  int
}

func exportFileCSV(fileName string, data [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	} else {
		write := csv.NewWriter(file)
		defer write.Flush()

		for _, value := range data {
			err := write.Write(value)
			if err != nil {
				log.Fatal("Cannot write file", err)
			}
		}
	}
}

func exportFile(fileName string, data []PersonalInfo) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	} else {
		write := csv.NewWriter(file)
		defer write.Flush()

		for _, value := range data {
			var arrData []string
			arrData = append(arrData, strconv.Itoa(value.id), value.name, strconv.Itoa(value.age))
			err := write.Write(arrData)
			if err != nil {
				log.Fatal("Cannot write file", err)
			}
		}
	}
}

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
