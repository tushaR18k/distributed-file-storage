package router

import (
	"api-gateway/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/upload", handlers.UploadFile).Methods("POST")
	r.HandleFunc("/api/download/{filename}", handlers.DownloadFile).Methods("GET")
	r.HandleFunc("/api/files", handlers.ListFiles).Methods("GET")

	return r
}
