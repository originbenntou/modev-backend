package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Conn *sqlx.DB
}

func NewDB(conn *sqlx.DB) *DB {
	return &DB{
		conn,
	}
}
