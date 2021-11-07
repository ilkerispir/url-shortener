package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	spreadsheetId := os.Getenv("spreadsheetId")
}
