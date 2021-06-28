package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func allBooks(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Books */
	var books []Book
	var reservations []Reservation

	database.Find(&books)
	
	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(books)

}

func avaibleBooks(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Books */
	var books []Book
	var reservations []Reservation

	database.Where("Stock > 0").Find(&books)
	
	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Book */
	vars := mux.Vars(r)
	id := vars["Id"]

	var book Book
	var reservations []Reservation

	database.First(&book, id)
	database.Model(&book).Related(&reservations)
	
	book.Reservations = reservations

	json.NewEncoder(w).Encode(book)
}

func newBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Create New Book */
	vars := mux.Vars(r)
	title := vars["Title"]
	stock := vars["Stock"]
	intStock, _ := strconv.Atoi(stock)

	database.Create(&Book{Title: title, Stock: intStock})

	fmt.Fprintf(w, "New Book Succesfuly Created!")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Delete Book */
	vars := mux.Vars(r)
	id := vars["Id"]

	var book Book
	database.Where("ID = ?", id).Find(&book)
	database.Delete(&book)

	fmt.Fprintf(w, "Book Succesfuly Deleted!")
}

func updateBookStock(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Update Book */
	vars := mux.Vars(r)
	id := vars["Id"]
	updatedStock := vars["Stock"]
	intUpdatedStock, _ := strconv.Atoi(updatedStock)

	var book Book
	database.Where("ID = ?", id).Find(&book)
	
	book.Stock = intUpdatedStock

	database.Save(&book)
	fmt.Fprintf(w, "Book Stock Succesfuly Updated!")
}