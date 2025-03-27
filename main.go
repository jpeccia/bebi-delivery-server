package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jpeccia/bebi-delivery-server/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
}
