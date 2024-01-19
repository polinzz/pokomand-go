package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "database:5657",
		DBName:               "gobase",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		http.HandleFunc("/erreur1", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Erreur 1 !"))
		})
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		http.HandleFunc("/erreur2", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Erreur 2 !"))
		})
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world !"))
	})

	http.ListenAndServe(":5686", nil)
}
