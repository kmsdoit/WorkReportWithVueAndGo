package main

import (
	"github.com/joho/godotenv"
	"github.com/kmsdoit/WorkReportWithVueAndGo/backend/routes"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//var db = utility.GetConnection()
	log.Println("Listening on Port 8081")
	routes.Router()
}
