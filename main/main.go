package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pokomand-go/Entity"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	router.Post("/user", Entity.GetUsers())
	router.Put("/user/add", Entity.AddUser())
	router.Get("/user/{id}", Entity.ShowUser())
	router.Post("/hub", Entity.GetHubs())

	// http.HandleFunc("/", Entity.GetUsers())

	http.ListenAndServe(":5686", router)
}
