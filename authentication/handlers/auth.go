package handlers

import (
	"authentication/middlewares"
	"authentication/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	app        *App
}

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddress string, app App) *APIServer {
	return &APIServer{
		listenAddr: listenAddress,
		app:        &app,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", makeHTTPHandleFunc(s.app.SignupHandler)).Methods("POST")
	router.HandleFunc("/login", makeHTTPHandleFunc(s.app.LoginHandler)).Methods("POST")
	router.HandleFunc("/logout", makeHTTPHandleFunc(s.app.LogoutHandler)).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middlewares.JWTMiddleware)

	log.Println("Running the server on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
