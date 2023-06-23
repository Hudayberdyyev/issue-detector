package repository

import "fmt"

type Reader struct {
	hostsDB, usersDB *Database
}

func NewReader(hostsDB, usersDB *Database) IReader {
	return &Reader{
		hostsDB: hostsDB,
		usersDB: usersDB,
	}
}

func (r *Reader) GetAccessStatusFromHosts(ip string) (int, error) {
	query := fmt.Sprintf(`SELECT active FROM hosts WHERE ip = $1`)
	var active int
	err := r.hostsDB.client.QueryRow(query, ip).Scan(&active)
	return active, err
}

func (r *Reader) GetAccessStatusFromUsersDB(userId int) (int, error) {
	query := fmt.Sprintf(`SELECT active FROM users_access_history WHERE id = $1`)
	var active int
	err := r.usersDB.client.QueryRow(query, userId).Scan(&active)
	return active, err
}
