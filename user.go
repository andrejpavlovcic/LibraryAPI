package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB
var err error

/* User Table */
type User struct {
	gorm.Model

	Name         string
	Surname      string
	Reservations []Reservation
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()
	
	var users []User
	database.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User")
}