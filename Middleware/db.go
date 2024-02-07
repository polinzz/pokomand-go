package Middleware

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
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

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))

	return fmt.Sprintf("%x", hash)
}
