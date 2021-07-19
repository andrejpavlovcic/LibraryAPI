package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//database global
var DB *gorm.DB

var databaseURI = "host=ec2-54-247-158-179.eu-west-1.compute.amazonaws.com user=nhncmcoribklwj dbname=d13gif6br221hd password=498ca2245aa1ef6280c2b5ee942e2cc974d333b435c3bd05629e94b0855ebb02 port=5432"

func SetupDB() *gorm.DB {

	//db config vars
	var dbHost string = os.Getenv("DB_HOST")
	var dbName string = os.Getenv("DB_NAME")
	var dbUser string = os.Getenv("DB_USERNAME")
	var dbPassword string = os.Getenv("DB_PASSWORD")
	var dbl = os.Getenv("DATABASE_URL")

	fmt.Println(dbHost)
	fmt.Println(dbName)
	fmt.Println(dbUser)
	fmt.Println(dbPassword)
	fmt.Println(dbl)

	//connect to db
	db, dbError := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if dbError != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection To DB Succesful!")

	//fix for connection timeout
	//see: https://github.com/go-sql-driver/mysql/issues/257
	db.DB().SetMaxIdleConns(0)

	return db
}
