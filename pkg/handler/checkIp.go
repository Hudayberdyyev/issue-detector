package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) checkIP(w http.ResponseWriter, r *http.Request) {
	// Simple authorization middleware
	if r.Header.Get(AuthorizationHeader) != h.secretKey {
		http.Error(w, "invalid secret key", http.StatusBadRequest)
		log.Printf("invalid secret key, incoming header = %s\n", h.secretKey)
		return
	}

	// Read the request body
	var checkIpRequest CheckIpBody
	err := json.NewDecoder(r.Body).Decode(&checkIpRequest)
	if err != nil {
		log.Printf("error when parse check ip request body: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If access allowed just log it and return success response
	switch checkIpRequest.ResponseIsAccessAllowed {
	case AccessGranted:
		// Just log it
		err := h.repo.IWriter.SaveCheckIpLogToDB(checkIpRequest.GetCheckIpModel())
		if err != nil {
			log.Printf("error when save check ip request log to DB: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case AccessDenied:
		// Compare it from with access on db
		userId := checkIpRequest.GetUserIdFromAccessToken()
		userAccess, err := h.repo.IReader.GetAccessStatusFromUsersDB(userId)
		if err != nil {
			log.Printf("%s\n", err.Error())
			return
		}
		ipAccess, err := h.repo.IReader.GetAccessStatusFromHosts(checkIpRequest.XForwardedForHeader)
		if err != nil {
			log.Printf("%s\n", err.Error())
		}
		err = h.repo.IWriter.SaveCheckIpLogToDB(checkIpRequest.GetCheckIpModelWithArgs(ipAccess, userAccess))
		if err != nil {
			log.Printf("error when save check ip request log to DB: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		log.Printf("checkIp access status value equal to = %d\n", checkIpRequest.ResponseIsAccessAllowed)

	}

	// Send a response
	w.WriteHeader(http.StatusOK)
}
