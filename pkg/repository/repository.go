package repository

type IReader interface {
}

type IWriter interface {
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
