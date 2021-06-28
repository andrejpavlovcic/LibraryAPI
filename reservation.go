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

	/* Find Reservations */
	var reservations []Reservation
	if err := database.Find(&reservations).Error; err != nil {
		fmt.Println(err.Error())
		panic("Reservation Find Error")
	}

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
	if err := database.First(&book, bookID).Error; err != nil {
		fmt.Println(err.Error())
		panic("Book Find Error")
	}

	if book.Stock <= 0 {
		fmt.Fprintf(w, "This Book Is Out Of Stock!")
	} else {
		/* Create Reservation */
		if err := database.Create(&Reservation{UserID: intUserID, BookID: intBookID}).Error; err != nil {
			fmt.Println(err.Error())
			panic("Reservation Create Error")
		}

		/* Update Book Stock */
		book.Stock = (book.Stock - 1)
		if err := database.Save(&book).Error; err != nil {
			fmt.Println(err.Error())
			panic("Book Update Error")
		}
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
	if err := database.Where("User_ID = ? AND Book_ID = ?", userID, bookID).Find(&reservation).Error; err != nil {
		fmt.Println(err.Error())
		panic("Reservation Find Error")
	}
	fmt.Printf("%+v\n", reservation)
	if err := database.Delete(&reservation).Error; err != nil {
		fmt.Println(err.Error())
		panic("Reservation Delete Error")
	}

	fmt.Fprintf(w, "Reservation Succesfuly Deleted!")
}