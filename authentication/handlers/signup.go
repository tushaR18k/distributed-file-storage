package handlers

import (
	"authentication/db"
	"authentication/models"
	"authentication/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type App struct {
	DB *db.PostgresStore
}

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewApp(db db.PostgresStore) *App {
	return &App{
		DB: &db,
	}
}

func (s *App) SignupHandler(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(SignUpRequest)
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createAccountReq.Password),
		bcrypt.DefaultCost)
	if err != nil {
		//http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return utils.WriteJSON(w, http.StatusInternalServerError, ApiError{Error: "Error hashing password"})
	}

	user := &models.User{
		Username: createAccountReq.Username,
		Email:    createAccountReq.Email,
		Password: string(hashedPassword),
	}

	if err := s.DB.StoreUser(*user); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusCreated, user)
}
