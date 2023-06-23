package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) checkIP(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var checkIpRequest CheckIpBody
	err := json.NewDecoder(r.Body).Decode(&checkIpRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If access allowed just log it and return success response
	switch checkIpRequest.ResponseIsAccessAllowed {
	case AccessGranted:
		// Just log it
		err := h.repo.IWriter.SaveCheckIpLogToDB(checkIpRequest.GetCheckIpModel())
		if err != nil {
			log.Println(err.Error())
		}
	case AccessDenied:
		// Compare it from with access on db
		userId := checkIpRequest.GetUserIdFromAccessToken()
		userAccess, err := h.repo.IReader.GetAccessStatusFromUsersDB(userId)
		if err != nil {
			log.Println(err.Error())
		}
		ipAccess, err := h.repo.IReader.GetAccessStatusFromHosts(checkIpRequest.XForwardedForHeader)
		if err != nil {
			log.Println(err.Error())
		}
		err = h.repo.IWriter.SaveCheckIpLogToDB(checkIpRequest.GetCheckIpModelWithArgs(ipAccess, userAccess))
		if err != nil {
			log.Println(err.Error())
		}
	default:
		log.Printf("checkIp access status value equal to = %d\n", checkIpRequest.ResponseIsAccessAllowed)

	}

	// Send a response
	w.WriteHeader(http.StatusOK)
}
