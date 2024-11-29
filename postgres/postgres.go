package postgres

import (
	"database/sql"
)

type ConnectDB struct {
	db *sql.DB
}

func NewConnectDB(pathConfig string) *ConnectDB {
	config, err := LoadConfig(pathConfig)
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
