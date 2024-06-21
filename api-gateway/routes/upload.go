package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the file on the server
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error creating file on server", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the file data to the server file
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying file data", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	log.Print("Encoding the names into JSON")

	log.Print("Sending back the response")
	w.WriteHeader(http.StatusOK)
	// Set the content-type header
	w.Header().Set("Content-Type", "application/json")

}
