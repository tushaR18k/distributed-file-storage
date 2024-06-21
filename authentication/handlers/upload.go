package handlers

import (
	"authentication/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (s *App) UploadHandler(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.UserIdFromContext(r.Context())
	if !ok {
		//http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User ID not found in context"})
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": "Error parsing form data"})
	}

	api_host := os.Getenv("API_HOST")
	api_port := os.Getenv("API_PORT")
	if api_host == "" {
		log.Fatal("Environment variable API_HOST is not set")
	}
	if api_port == "" {
		log.Fatal("Environment variable API_PORT is not set")
	}

	log.Print("Starting uploading file!")

	body, err, headerType := prepareRequestBody(w, r)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s:%s/api/upload", api_host, api_port)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to create request"})
	}
	req.Header.Set("Content-Type", headerType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to call api-gateway /upload endpoint"})
	}
	defer resp.Body.Close()

	// Copy the headers from the api-gateway response to the client response
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Copy the response body from api-gateway to the client
	if _, err := io.Copy(w, resp.Body); err != nil {
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Failed to write response body api-gateway /upload endpoint"})
	}

	return nil
}

func (s *App) UserIdFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value("userID").(uint)
	return userID, ok
}

func prepareRequestBody(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, error, string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, values := range r.MultipartForm.Value {
		for _, value := range values {
			_ = writer.WriteField(key, value)
		}
	}

	for key, files := range r.MultipartForm.File {
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				return nil, utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error opening file"}), ""
			}
			defer file.Close()

			part, err := writer.CreateFormFile(key, fileHeader.Filename)
			if err != nil {
				return nil, utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error creating form file"}), ""
			}

			_, err = io.Copy(part, file)
			if err != nil {
				return nil, utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error copying file"}), ""
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error closing writer"}), ""
	}

	return body, nil, writer.FormDataContentType()
}
