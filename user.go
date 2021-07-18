package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	/* Find Users */
	var users []User
	var reservations []Reservation

	if err := database.Find(&users).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Find Error")
	}

	for index := range users {
		database.Model(&users[index]).Related(&reservations)
		users[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(users)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find User */
	vars := mux.Vars(r)
	id := vars["Id"]

	var user User
	var reservations []Reservation

	if err := database.First(&user, id).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Find Error")
	}
	database.Model(&user).Related(&reservations)

	user.Reservations = reservations

	json.NewEncoder(w).Encode(user)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Create New User */
	vars := mux.Vars(r)
	name := vars["Name"]
	surname := vars["Surname"]

	if err := database.Create(&User{Name: name, Surname: surname}).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Create Error")
	}

	fmt.Fprintf(w, "New User Succesfuly Created!")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Delete User */
	vars := mux.Vars(r)
	id := vars["Id"]

	var user User
	if err := database.Where("ID = ?", id).Find(&user).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Delete Error")
	}

	if err := database.Unscoped().Delete(&user).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Delete Error")
	}

	fmt.Fprintf(w, "User Succesfuly Deleted!")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Update User */
	vars := mux.Vars(r)
	id := vars["Id"]
	newName := vars["NewName"]
	newSurname := vars["NewSurname"]

	var user User
	if err := database.Where("ID = ?", id).Find(&user).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Update Error - No User")
	}

	user.Name = newName
	user.Surname = newSurname

	if err := database.Save(&user).Error; err != nil {
		fmt.Println(err.Error())
		panic("User Update Error")
	}
	fmt.Fprintf(w, "User Succesfuly Updated!")
}
