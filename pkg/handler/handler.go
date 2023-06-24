package handler

import (
	"issue-detector/pkg/repository"
	"net/http"
)

type Handler struct {
	secretKey string
	repo      *repository.Repository
}

func NewHandler(repo *repository.Repository, secretKey string) *Handler {
	return &Handler{
		secretKey: secretKey,
		repo:      repo,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method + ":" + req.URL.Path {
	case "POST:/checkIp":
		h.checkIP(w, req)
	case "POST:/refresh_token":
		h.refreshToken(w, req)
	default:
		http.NotFound(w, req)
	}
}
