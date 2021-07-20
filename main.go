package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Andre711/LibraryAPI/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//Init Router
	port := os.Getenv("PORT")
	router := NewRouter()

	//Setup Database
	db.DB = db.SetupDB()
	defer db.DB.Close()

	//Create Http Server
	log.Fatal(http.ListenAndServe(":"+port, router))

}
