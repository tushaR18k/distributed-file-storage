package main

import (
	"api-gateway/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	fmt.Println("Server is running on http:localhost:8000")
	if err := http.ListenAndServe(":8000", corsHandler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
