package handler

import (
	"net/http"
)

func checkIP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
