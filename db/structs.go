package db

import "github.com/jinzhu/gorm"

//Book Table Struct
type Book struct {
	gorm.Model

	Title        string
	Stock        int
	Reservations []Reservation
}

//User Table Struct
type User struct {
	gorm.Model

	Name         string
	Surname      string
	Reservations []Reservation
}

//Reservation Table Struct
type Reservation struct {
	gorm.Model

	UserID int
	BookID int
}
