package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func allBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book
	var reservations []Reservation

	if err := database.Find(&books).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func avaibleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book
	var reservations []Reservation

	if err := database.Where("Stock > 0").Find(&books).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["ID"]

	var book Book
	var reservations []Reservation

	if err := database.First(&book, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	database.Model(&book).Related(&reservations)
	book.Reservations = reservations

	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["ID"]

	var book Book
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := database.Unscoped().Delete(&book).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func updateBookStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["ID"]
	updatedStock := vars["Stock"]
	intUpdatedStock, _ := strconv.Atoi(updatedStock)

	var book Book
	if err := database.Where("ID = ?", id).Find(&book).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	book.Stock = intUpdatedStock
	database.Save(&book)

	if err := database.Save(&book).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
