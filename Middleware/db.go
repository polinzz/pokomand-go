package Middleware

import (
	"database/sql"
	"log"
)

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(database:3306)/gobase")
	if err != nil {
		log.Fatal("err1 ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("err2 ", err)

	}

	return db
}
