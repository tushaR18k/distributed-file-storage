package routes

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Get the filename from the URL path
	filename := r.URL.Path[len("/api/download/"):]

	// Open the file on the server
	f, err := os.Open("./uploads/" + filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	// Set the appropriate content-type header
	w.Header().Set("Content-Type", "application/octet-stream")

	// Copy the file data to the response writer
	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "Error copying file data", http.StatusInternalServerError)
		return
	}

}
