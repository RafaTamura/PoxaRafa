package main

import (
	"log"
	"os"
	"poxa_rafa/database"
	"poxa_rafa/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: não foi possível carregar o arquivo .env")
	}
}

func main() {
	database.Connect()
	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "PORT"
	}

	router.Run("0.0.0.0:" + port)
}
