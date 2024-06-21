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
	port := ":8000"

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	fmt.Printf("Server is running on http:localhost%s", port)
	if err := http.ListenAndServe(port, corsHandler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
