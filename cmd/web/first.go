package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

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
