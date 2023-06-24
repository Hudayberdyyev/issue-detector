package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	// Simple authorization middleware
	if r.Header.Get(AuthorizationHeader) != h.secretKey {
		http.Error(w, "invalid secret key", http.StatusBadRequest)
		return
	}

	// Read the request body
	var refreshTokenRequest RefreshTokenBody
	err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.IWriter.SaveRefreshTokenToDB(refreshTokenRequest.GetRefreshTokenDbModel()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
}
