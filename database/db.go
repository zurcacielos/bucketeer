package database

import (
	"database/sql"
	"fmt"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {

}

func Initialize(username, password, database string) (Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return db
	}
	db := Database{}
	return db, nil
}