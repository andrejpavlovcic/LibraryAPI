package main

import (
	"fmt"
	"log"
	"net/http"

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

/*
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	database.Find(users)

	fmt.Println("All Users Endpoint")
	json.NewEncoder(w).Encode(users)
}
*/

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Library API")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	//http.HandleFunc("/users", getUsers)
	log.Fatal(http.ListenAndServe(":5432", nil))
}

func main() {
	/* Database URI */
	databasebURI := "host=ec2-54-247-158-179.eu-west-1.compute.amazonaws.com user=nhncmcoribklwj dbname=d13gif6br221hd password=498ca2245aa1ef6280c2b5ee942e2cc974d333b435c3bd05629e94b0855ebb02 port=5432"

	/* Open Connection To Database */
	database, error := gorm.Open("postgres", databasebURI)
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("Sucessfully Connected To Database!")
	}
	/* Close Connection To Database */
	defer database.Close()

}
