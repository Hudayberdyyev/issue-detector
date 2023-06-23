package repository

type Writer struct {
	checkIpDB *Database
}

func NewWriter(checkIpDB *Database) IWriter {
	return &Writer{checkIpDB: checkIpDB}
}
