package user

import (
	"encoding/json"
	"net/http"

	"github.com/Andre711/LibraryAPI/db"
	customResponse "github.com/Andre711/LibraryAPI/response"
	"github.com/gorilla/mux"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []db.User
	var reservations []db.Reservation

	if err := db.DB.Find(&users).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, "no users found")
		return
	}

	for index := range users {
		db.DB.Model(&users[index]).Related(&reservations)
		users[index].Reservations = reservations
	}

	json.NewEncoder(w).Encode(users)

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["ID"]

	var user db.User
	var reservations []db.Reservation

	if err := db.DB.First(&user, id).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	db.DB.Model(&user).Related(&reservations)

	user.Reservations = reservations

	json.NewEncoder(w).Encode(user)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	name := vars["Name"]
	surname := vars["Surname"]

	var user db.User

	user.Name = name
	user.Surname = surname

	if err := db.DB.Create(&user).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, "couldn't create user")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["ID"]

	var user db.User
	if err := db.DB.Where("ID = ?", id).Find(&user).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if err := db.DB.Unscoped().Delete(&user).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, "couldn't delete user")
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["ID"]
	newName := vars["NewName"]
	newSurname := vars["NewSurname"]

	var user db.User
	if err := db.DB.Where("ID = ?", id).Find(&user).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	user.Name = newName
	user.Surname = newSurname

	if err := db.DB.Save(&user).Error; err != nil {
		customResponse.NewErrorResponse(w, http.StatusNotFound, "couldn't update user")
		return
	}

	json.NewEncoder(w).Encode(user)
}
