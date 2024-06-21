package handlers

import (
	"authentication/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (s *App) DownloadHandler(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.UserIdFromContext(r.Context())
	if !ok {
		//http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User ID not found in context"})
	}
	filename := r.URL.Path[len("/api/download/"):]
	api_host := os.Getenv("API_HOST")
	api_port := os.Getenv("API_PORT")
	if api_host == "" {
		log.Fatal("Environment variable API_HOST is not set")
	}
	if api_port == "" {
		log.Fatal("Environment variable API_PORT is not set")
	}
	url := fmt.Sprintf("http://%s:%s/api/download/%s", api_host, api_port, filename)
	log.Print("Making request to get the file: ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create request"})
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to call api-gateway /download endpoint"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to download file api-gateway /download endpoint"})
	}

	//Copying the headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(http.StatusOK)

	if _, err := io.Copy(w, resp.Body); err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to write response body for file api-gateway /download endpoint"})
	}
	log.Print("Successfully sent the file")
	return nil

}
