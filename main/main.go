package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pokomand-go/Entity"
	Store "pokomand-go/store"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	router.Post("/user", Entity.GetAllUsers())
	router.Put("/user/add", Entity.AddUser())
	router.Get("/user/{id}", Entity.GetUserById())
	router.Post("/login", Store.Login())
	router.Post("/hub", Entity.GetHubs())

	// http.HandleFunc("/", Entity.GetUsers())

	http.ListenAndServe(":5686", router)
}
