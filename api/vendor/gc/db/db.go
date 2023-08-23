package db

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type DBConf struct {
	User string
	Pass string
	Name string
	Host string
}

var dbConfig = DBConf{
	User: "gcuser",
	Pass: "gcpass",
	Name: "gc",
	Host: "localhost",
}

// generate the database connection string
func GenerateConnStr(c DBConf) string {
	s := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Name,
	)

	return s
}

// connect to the database
func Connect() (db *sql.DB, err error) {
	// generate the database connection string
	s := GenerateConnStr(dbConfig)

	// connect to the database
	db, err = sql.Open("postgres", s)
	if err != nil {
		return nil, err
	}

	// return the database connection
	return db, nil
}
