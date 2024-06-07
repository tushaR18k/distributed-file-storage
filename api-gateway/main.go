package main

import (
	"api-gateway/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on http:localhost:8000")
	http.ListenAndServe(":8000", r)
}
