package infra

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://user:pass@host:5432/db")
	if err != nil {
		panic(err)
	}
}
