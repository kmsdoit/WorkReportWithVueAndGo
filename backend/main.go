package main

import (
	"github.com/joho/godotenv"
	"github.com/kmsdoit/WorkReportWithVueAndGo/backend/routes"
	user "github.com/kmsdoit/WorkReportWithVueAndGo/backend/services/User"
	"github.com/kmsdoit/WorkReportWithVueAndGo/backend/utility"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var db = utility.GetConnection()
	user.SetDB(db)
	log.Println("Listening on Port 8082")
	routes.Router()
}
