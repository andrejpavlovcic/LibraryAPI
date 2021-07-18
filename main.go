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

var database *gorm.DB
var err error

var databaseURI = "host=ec2-54-247-158-179.eu-west-1.compute.amazonaws.com user=nhncmcoribklwj dbname=d13gif6br221hd password=498ca2245aa1ef6280c2b5ee942e2cc974d333b435c3bd05629e94b0855ebb02 port=5432"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Library API")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")

	/* Users */
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{Id}", getUser).Methods("GET")
	myRouter.HandleFunc("/user/{Name}/{Surname}", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{Id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{Id}/{NewName}/{NewSurname}", updateUser).Methods("PUT")

	/* Books */
	myRouter.HandleFunc("/books/all", allBooks).Methods("GET")
	myRouter.HandleFunc("/books/avaible", avaibleBooks).Methods("GET")
	myRouter.HandleFunc("/book/{Id}", getBook).Methods("GET")
	myRouter.HandleFunc("/book/{Title}/{Stock}", newBook).Methods("POST")
	myRouter.HandleFunc("/book/{Id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{Id}/{Stock}", updateBookStock).Methods("PUT")

	/* Reservations */
	myRouter.HandleFunc("/reservations", allReservations).Methods("GET")
	myRouter.HandleFunc("/reservation/{UserID}/{BookID}", newReservation).Methods("POST")
	myRouter.HandleFunc("/reservation/{UserID}/{BookID}", deleteReservation).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), myRouter))
}

func main() {
	_, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	handleRequests()

	//defer database.Close()
}
