package handlers

import (
	"authentication/models"
	"authentication/utils"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *App) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	loginReq := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}
	var user *models.User
	user, err := s.DB.FindUser(loginReq.Username)
	if err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: "Invalid username or password"})
	}

	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Login Successful!",
		"token": tokenString})
}
