package repository

import "fmt"

type Writer struct {
	checkIpDB *Database
}

func NewWriter(checkIpDB *Database) IWriter {
	return &Writer{checkIpDB: checkIpDB}
}

func (w *Writer) SaveCheckIpLogToDB(checkIpData CheckIpModel) error {
	query := fmt.Sprintf(`INSERT INTO check_ip_history(
                             authorization_header, user_agent_header, 
                             mac_address_user_header, x_forwarded_for_header, 
                             response_status_code, response_is_actual_version, 
                             response_is_access_allowed, db_is_ip_access_allowed, 
                             db_is_user_access_allowed, error_log, error_code)
						     VALUES (:authorization_header, :user_agent_header, 
						         :mac_address_user_header, :x_forwarded_for_header, 
						         :response_status_code, :response_is_actual_version,
						         :response_is_access_allowed, :db_is_ip_access_allowed,
						         :db_is_user_access_allowed, :error_log, :error_code)`)
	_, err := w.checkIpDB.client.NamedExec(query, checkIpData)
	return err
}
