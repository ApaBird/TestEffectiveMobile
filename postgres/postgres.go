package postgres

import (
	"database/sql"
)

type ConnectDB struct {
	db *sql.DB
}

func NewConnectDB() (*ConnectDB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", prepareConnectionString(config))
	if err != nil {
		return nil, err
	}

	return &ConnectDB{db: db}, nil
}

func prepareConnectionString(config *DBConfig) string {
	return "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DBName + " sslmode=" + config.Sslmode
}

func (c *ConnectDB) Close() error {
	return c.db.Close()
}
