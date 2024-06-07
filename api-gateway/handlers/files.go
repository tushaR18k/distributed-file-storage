package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	// Read the contents of the uploads directory
	files, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Error reading files", http.StatusInternalServerError)
		return
	}

	// Create a slice to store file names
	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	// Encode the file names slice to JSON
	jsonData, err := json.Marshal(filenames)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the content-type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)

}
