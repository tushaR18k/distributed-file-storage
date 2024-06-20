package handlers

import (
	"authentication/utils"
	"net/http"
)

func (s *App) FilesHandler(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.UserIdFromContext(r.Context())
	if !ok {
		//http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User ID not found in context"})
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "success!"})

}
