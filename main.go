package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/* User Table */
type User struct {
	gorm.Model

	Name         string
	Surname      string
	Reservations []Reservation
}

/* Book Table */
type Book struct {
	gorm.Model

	Title        string
	Stock        int
	Reservations []Reservation
}

/* Reservation Table */
type Reservation struct {
	gorm.Model

	UserID int
	BookID int
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Library API")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), myRouter))
}

func main() {
	fmt.Println("Library API")

	handleRequests()
}