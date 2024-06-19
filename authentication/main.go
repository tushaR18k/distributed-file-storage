package main

import (
	database "authentication/db"
	"authentication/handlers"
	"authentication/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := database.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	storage.DB.AutoMigrate(&models.User{})
	app := handlers.NewApp(*storage)
	router := handlers.NewAPIServer(":7000", *app)
	router.Run()
	defer storage.DB.Close()

}
