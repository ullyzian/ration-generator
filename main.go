package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"ration-generator/controllers"
	"ration-generator/db"
)


func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	conn := db.Connect()
	db.CreateTable(conn)

	// file server
	fs := http.FileServer(http.Dir("./static"))

	// handlers
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", controllers.RootHandler)
	http.HandleFunc("/programs/", controllers.ProgramsHandler)
	http.HandleFunc("/dishes/create/", controllers.CreateDish)
	http.HandleFunc("/dishes/", controllers.GetDishes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

