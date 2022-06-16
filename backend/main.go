package main

import (
	"github.com/joho/godotenv"
	"github.com/kmsdoit/WorkReportWithVueAndGo/backend/routes"
	"github.com/kmsdoit/WorkReportWithVueAndGo/backend/utility"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go utility.DBConnection()

	log.Println("Listening on Port 8082")
	routes.Router()
}
