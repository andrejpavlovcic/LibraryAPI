package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Database Global
var DB *gorm.DB

func SetupDB() *gorm.DB {

	//Connect To DB
	db, dbError := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if dbError != nil {
		panic("Failed To Connect To Database!")
	}

	fmt.Println("Connection To DB Succesful!")

	db.DB().SetMaxIdleConns(0)

	return db
}
