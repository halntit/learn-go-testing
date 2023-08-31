package db

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

type PostgresConn struct {
	DB *sql.DB
}
