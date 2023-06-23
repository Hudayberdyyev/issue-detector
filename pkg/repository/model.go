package repository

type CheckIpModel struct {
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
