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

/* Reservation Table */
type Reservation struct {
	gorm.Model

	UserID int
	BookID int
}

func allReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/* Find Reservations */
	var reservations []Reservation

	if err := db.DB.Find(&reservations).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(reservations)
}

func newReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["UserID"]
	bookID := vars["BookID"]
	intUserID, _ := strconv.Atoi(userID)
	intBookID, _ := strconv.Atoi(bookID)

	var book Book
	var reservation Reservation

	if err := db.DB.First(&book, bookID).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if book.Stock <= 0 {
		customResponse.NewErrorResponse(w, http.StatusNotFound, "Out Of Stock")
		return
	}

	reservation.UserID = intUserID
	reservation.BookID = intBookID

	/* Create Reservation */
	if err := db.DB.Create(&reservation).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	/* Update Book Stock */
	book.Stock = (book.Stock - 1)

	if err := db.DB.Save(&book).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(reservation)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	/* Delete Reservation */
	vars := mux.Vars(r)
	userID := vars["UserID"]
	bookID := vars["BookID"]

	var reservation Reservation

	if err := db.DB.Where("User_ID = ? AND Book_ID = ?", userID, bookID).Find(&reservation).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if err := db.DB.Unscoped().Delete(&reservation).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(w).Encode(reservation)
}
