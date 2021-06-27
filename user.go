package main

import (
	"fmt"
	"net/http"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users")
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