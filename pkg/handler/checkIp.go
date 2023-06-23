package handler

import (
	"encoding/json"
	"net/http"
)

type CheckIpBody struct {
	AuthorizationHeader     string `json:"authorization_header"`
	UserAgentHeader         string `json:"user_agent_header"`
	MacAddressHeader        string `json:"mac_address_user_header"`
	XForwardedForHeader     string `json:"x_forwarded_for_header"`
	ResponseStatusCode      int    `json:"response_status_code"`
	ResponseIsActualVersion int    `json:"response_is_actual_version"`
	ResponseIsAccessAllowed int    `json:"response_is_access_allowed"`
	ErrorLog                string `json:"error_log"`
	ErrorCode               string `json:"error_code"`
}

func checkIP(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var checkIpRequest CheckIpBody
	err := json.NewDecoder(r.Body).Decode(&checkIpRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
}
