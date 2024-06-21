package handlers

import (
	"authentication/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (s *App) FilesHandler(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.UserIdFromContext(r.Context())
	if !ok {
		return utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": "User ID not found in context"})
	}
	api_host := os.Getenv("API_HOST")
	api_port := os.Getenv("API_PORT")
	if api_host == "" {
		log.Fatal("Environment variable API_HOST is not set")
	}
	if api_port == "" {
		log.Fatal("Environment variable API_PORT is not set")
	}
	url := fmt.Sprintf("http://%s:%s/api/files", api_host, api_port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create request"})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to call api-gateway /files endpoint"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to get files api-gateway /files endpoint"})
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to read response body files api-gateway /files endpoint"})
	}

	// Forward the response from the api-gateway to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return nil

}
