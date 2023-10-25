package db

import (
	"database/sql"
	"fmt"

	"github.com/gc/types"

	_ "github.com/lib/pq"
)

var dbConfig = types.DBConf{
	User: "gcuser",
	Pass: "gcpass",
	Name: "gc",
	Host: "localhost",
}

// generate the database connection string
func GenerateConnStr(c types.DBConf) string {
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
