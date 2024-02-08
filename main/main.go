package main

import (
	"net/http"
	"pokomand-go/Entity"
	Store "pokomand-go/store"

	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	router.Get("/user", Entity.GetAllUsers())
	router.Put("/user/add", Store.SignUp())
	// router.Get("/user/{id}", Entity.GetUserById())
	router.Post("/login", Store.Login())
	router.Get("/hub", Entity.GetAllHubs())
	// router.Post("/hub/add", Entity.AddHub())

	// http.HandleFunc("/", Entity.GetUsers())

	http.ListenAndServe(":5686", router)
}
