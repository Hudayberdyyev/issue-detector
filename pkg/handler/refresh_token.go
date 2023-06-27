package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	// Simple authorization middleware
	if r.Header.Get(AuthorizationHeader) != h.secretKey {
		log.Printf("authorization header doesn't match, incoming header = %s\n", h.secretKey)
		http.Error(w, "invalid secret key", http.StatusBadRequest)
		return
	}

	// Read the request body
	var refreshTokenRequest RefreshTokenBody
	err := json.NewDecoder(r.Body).Decode(&refreshTokenRequest)
	if err != nil {
		log.Printf("can't parse refresh token body: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.IWriter.SaveRefreshTokenToDB(refreshTokenRequest.GetRefreshTokenDbModel()); err != nil {
		log.Printf("error when save refresh token request log to DB: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
}
