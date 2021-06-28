package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func allBooks(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Books */
	var books []Book
	var reservations []Reservation

	database.Find(&books)
	
	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(books)

}

func avaibleBooks(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Books */
	var books []Book
	var reservations []Reservation

	database.Where("Stock > 0").Find(&books)
	
	for index := range books {
		database.Model(&books[index]).Related(&reservations)
		books[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Find Book */
	vars := mux.Vars(r)
	id := vars["Id"]

	var book Book
	var reservations []Reservation

	database.First(&book, id)
	database.Model(&book).Related(&reservations)
	
	book.Reservations = reservations

	json.NewEncoder(w).Encode(book)
}

func newBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Create New Book */
	vars := mux.Vars(r)
	name := vars["Name"]
	surname := vars["Surname"]

	database.Create(&User{Name: name, Surname: surname})

	fmt.Fprintf(w, "New User Succesfuly Created!")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Delete Book */
	vars := mux.Vars(r)
	id := vars["Id"]

	var user User
	database.Where("ID = ?", id).Find(&user)
	database.Delete(&user)

	fmt.Fprintf(w, "User Succesfuly Deleted!")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	database, err := gorm.Open("postgres", databaseURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect To Database!")
	}
	defer database.Close()

	/* Update Book */
	vars := mux.Vars(r)
	id := vars["Id"]
	newName := vars["NewName"]
	newSurname := vars["NewSurname"]

	var user User
	database.Where("ID = ?", id).Find(&user)
	
	user.Name = newName
	user.Surname = newSurname

	database.Save(&user)
	fmt.Fprintf(w, "User Succesfuly Updated!")
}