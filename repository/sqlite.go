package repository

import (
	"database/sql"
)

func NewSqliteDB(driverName string, pathDB string) (*sql.DB, error) {
	db, err := sql.Open(driverName, pathDB)
	return db, err
}
