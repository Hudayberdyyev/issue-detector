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
	switch req.Method + ":" + req.URL.Path {
	case "POST:/checkIp":
		checkIP(w, req)
	default:
		http.NotFound(w, req)
	}
}
