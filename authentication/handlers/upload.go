package handlers

import (
	"authentication/utils"
	"context"
	"net/http"
)

func (s *App) UploadHandler(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.UserIdFromContext(r.Context())
	if !ok {
		//http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User ID not found in context"})
	}

	return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "success!"})
}

func (s *App) UserIdFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value("userID").(uint)
	return userID, ok
}
