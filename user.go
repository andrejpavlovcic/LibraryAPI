package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	database, err := gorm.Open("postgres", "host=ec2-54-247-158-179.eu-west-1.compute.amazonaws.com user=nhncmcoribklwj dbname=d13gif6br221hd password=498ca2245aa1ef6280c2b5ee942e2cc974d333b435c3bd05629e94b0855ebb02 port=5432")
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