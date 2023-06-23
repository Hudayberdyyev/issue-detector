package repository

type Reader struct {
	hostsDB, usersDB *Database
}

func NewReader(hostsDB, usersDB *Database) IReader {
	return &Reader{
		hostsDB: hostsDB,
		usersDB: usersDB,
	}
}

func (r *Reader) GetAccessStatusFromHosts() (int, error) {
	return 0, nil
}

func (r *Reader) GetAccessStatusFromUsersDB() (int, error) {
	return 0, nil
}
