package repository

type RefreshTokenModel struct {
	UserId               int    `db:"user_id"`
	RequestFingerprint   string `db:"request_fingerprint"`
	RequestRefreshToken  string `db:"request_refresh_token"`
	ResponseRefreshToken string `db:"response_refresh_token"`
	ResponseAccessToken  string `db:"response_access_token"`
	ResponseStatusCode   int    `db:"response_status_code"`
	ResponseMessage      string `db:"response_message"`
	ErrorCode            string `db:"err_code"`
}

type CheckIpModel struct {
	UserId                  int    `db:"user_id"`
	AuthorizationHeader     string `db:"authorization_header"`
	UserAgentHeader         string `db:"user_agent_header"`
	MacAddressHeader        string `db:"mac_address_user_header"`
	XForwardedForHeader     string `db:"x_forwarded_for_header"`
	ResponseStatusCode      int    `db:"response_status_code"`
	ResponseIsActualVersion int    `db:"response_is_actual_version"`
	ResponseIsAccessAllowed int    `db:"response_is_access_allowed"`
	DbIsIpAccessAllowed     int    `db:"db_is_ip_access_allowed"`
	DbIsUserAccessAllowed   int    `db:"db_is_user_access_allowed"`
	ErrorLog                string `db:"error_log"`
	ErrorCode               string `db:"error_code"`
}
