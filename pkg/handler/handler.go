package handler

import (
	"issue-detector/pkg/repository"
	"net/http"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (r *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/checkIp":
		checkIP(w, req)
	default:
		http.NotFound(w, req)
	}
}
