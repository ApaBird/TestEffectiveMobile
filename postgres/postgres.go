package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type ConnectDB struct {
	db *sql.DB
}

func Init() {
	db, err := NewConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	files, err := os.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		return
	}

	migrations := make(map[int]string)
	rollback := make(map[int]string)
	for _, file := range files {
		name := strings.Split(file.Name(), "_")

		num, err := strconv.Atoi(name[0])
		if err != nil {
			panic(fmt.Errorf("invalid migration name: %s", file.Name()))
		}

		body, err := os.ReadFile("migrations/" + file.Name())
		if err != nil {
			panic(err)
		}

		if strings.Contains(name[1], "up") {
			migrations[num] = string(body)
		} else if strings.Contains(name[1], "down") {
			rollback[num] = string(body)
		} else {
			panic(fmt.Errorf("invalid migration name: %s", file.Name()))
		}

	}

	fmt.Println(migrations)

	for i := 0; i < len(migrations); i++ {
		migration := migrations[i]
		fmt.Println("Apply migration: " + migration)
		if _, err := db.db.Exec(migration); err != nil {
			for j := i; j >= 0; j-- {
				back := rollback[j]
				if _, err := db.db.Exec(back); err != nil {
					panic(err)
				}
			}
			panic(err)
		}
	}

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
