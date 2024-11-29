package postgres

import (
	"database/sql"
)

type ConnectDB struct {
	db *sql.DB
}

func NewConnectDB() *ConnectDB {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", prepareConnectionString(config))
	if err != nil {
		panic(err)
	}

	return &ConnectDB{db: db}
}

func prepareConnectionString(config *DBConfig) string {
	return "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DBName + " sslmode=" + config.Sslmode
}
