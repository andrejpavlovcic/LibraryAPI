package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Andre711/LibraryAPI/db"
	customResponse "github.com/Andre711/LibraryAPI/response"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

/* Book Table */
type Book struct {
	gorm.Model

	Title        string
	Stock        int
	Reservations []Reservation
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book
	var reservations []Reservation

	if err := db.DB.Find(&books).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	for index := range books {
		db.DB.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func avaibleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book
	var reservations []Reservation

	if err := db.DB.Where("Stock > 0").Find(&books).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	for index := range books {
		db.DB.Model(&books[index]).Related(&reservations)
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

	if err := db.DB.First(&book, id).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	db.DB.Model(&book).Related(&reservations)
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

	if err := db.DB.Create(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
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
	if err := db.DB.Where("ID = ?", id).Find(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if err := db.DB.Unscoped().Delete(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
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
	if err := db.DB.Where("ID = ?", id).Find(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	book.Stock = intUpdatedStock
	db.DB.Save(&book)

	if err := db.DB.Save(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
