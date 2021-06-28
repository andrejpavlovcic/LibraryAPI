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

	if err := database.Find(&books).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}
	
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

	if err := database.Where("Stock > 0").Find(&books).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}
	
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

	if err := database.First(&book, id).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}
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

	if err := database.Create(&Book{Title: title, Stock: intStock}).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Create Error")
	}

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
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}
	
	if err := database.Delete(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Delete Error")
	}

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
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}
	
	book.Stock = intUpdatedStock

	if err := database.Save(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Update Error")
	}
	fmt.Fprintf(w, "Book Stock Succesfuly Updated!")
}