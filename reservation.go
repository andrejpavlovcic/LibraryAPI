package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func allReservations(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Books */
	var reservations []Reservation
	database.Find(&reservations)

	json.NewEncoder(w).Encode(reservations)

}

func newReservation(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	vars := mux.Vars(r)
	userID := vars["UserID"]
	bookID := vars["BookID"]
	intUserID, _ := strconv.Atoi(userID)
	intBookID, _ := strconv.Atoi(bookID)

	var book Book
	database.First(&book, bookID)

	if book.Stock <= 0 {
		fmt.Fprintf(w, "This Book Is Out Of Stock!")
	} else {
		/* Create Reservation */
		database.Create(&Reservation{UserID: intUserID, BookID: intBookID})

		/* Update Book Stock */
		book.Stock = (book.Stock - 1)
		database.Save(&book)
		fmt.Fprintf(w, "Reservation Succesfuly Created!")
	}
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Delete Reservation */
	vars := mux.Vars(r)
	userID := vars["UserID"]
	bookID := vars["BookID"]

	var reservation Reservation
	database.Where("UserID = ? AND BookID = ?", userID, bookID).Find(&reservation)
	database.Delete(&reservation)

	fmt.Fprintf(w, "Reservation Succesfuly Deleted!")
}