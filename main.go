package main

import (
	"log"

	"example.com/config/db"
	"example.com/models"

	"example.com/routers"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.GetConnection()
	models.UserInit()
}

func main() {
	routers.InitRoute()
}
