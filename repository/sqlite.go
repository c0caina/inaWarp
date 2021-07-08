package repository

import (
	"database/sql"
)

var WarpRepo *WarpSqlite

func NewSqliteDB(driverName string, pathDB string) (*sql.DB, error) {
	db, err := sql.Open(driverName, pathDB)
	return db, err
}
