package repository

type IReader interface {
	GetAccessStatusFromUsersDB(userId int) (int, error)
	GetAccessStatusFromHosts(ip string) (int, error)
}

type IWriter interface {
	SaveCheckIpLogToDB(checkIpData CheckIpModel) error
	SaveRefreshTokenToDB(refreshTokenData RefreshTokenModel) error
}

type Repository struct {
	IReader
	IWriter
}

func NewRepository(hostsDB, usersDB, checkIpDB *Database) *Repository {
	return &Repository{
		IReader: NewReader(hostsDB, usersDB),
		IWriter: NewWriter(checkIpDB),
	}
}
