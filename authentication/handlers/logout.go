package handlers

import (
	"authentication/utils"
	"net/http"
	"time"
)

func (s *App) LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Successfully logged out!"})
}
