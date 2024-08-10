package database

var connection string

func init() {
	connection = "SQL Server"
}

func GetDatabate() string {
	return connection
}
