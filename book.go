package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func allBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* Find All Books */
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
	w.Header().Set("Content-Type", "application/json")

	/* Find All AvaibleBooks */
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
	w.Header().Set("Content-Type", "application/json")

	/* Find Book */
	vars := mux.Vars(r)
	id := vars["ID"]

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
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	vars := mux.Vars(r)
	title := vars["Title"]
	stock := vars["Stock"]
	intStock, _ := strconv.Atoi(stock)

	book.Title = title
	book.Stock = intStock

	if err := database.Create(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Create Error")
	}

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* Delete Book */
	vars := mux.Vars(r)
	id := vars["ID"]

	var book Book
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}

	if err := database.Unscoped().Delete(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Delete Error")
	}

	json.NewEncoder(w).Encode(book)
}

func updateBookStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* Update Book */
	vars := mux.Vars(r)
	id := vars["ID"]
	updatedStock := vars["Stock"]
	intUpdatedStock, _ := strconv.Atoi(updatedStock)

	var book Book
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		json.NewEncoder(w).Encode("")
	}

	book.Stock = intUpdatedStock

	if err := database.Save(&book).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Update Error")
	}

	json.NewEncoder(w).Encode(book)
}
