package router

import (
	"api-gateway/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/upload", routes.UploadFile).Methods("POST")
	r.HandleFunc("/api/download/{filename}", routes.DownloadFile).Methods("GET")
	r.HandleFunc("/api/files", routes.ListFiles).Methods("GET")

	return r
}
