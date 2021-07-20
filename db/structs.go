package db

import "github.com/jinzhu/gorm"

/* Book Struct */
type Book struct {
	gorm.Model

	Title        string
	Stock        int
	Reservations []Reservation
}

/* User Struct */
type User struct {
	gorm.Model

	Name         string
	Surname      string
	Reservations []Reservation
}

/* Reservation Struct */
type Reservation struct {
	gorm.Model

	UserID int
	BookID int
}
