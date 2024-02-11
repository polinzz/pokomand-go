package main

import (
	"net/http"
	"pokomand-go/Entity"
	Store "pokomand-go/store"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router.Use(cors.Handler)

	// ENDPOINTS

	// Users
	router.Get("/user", Entity.GetAllUsers())
	router.Post("/user/add", Store.SignUp())
	router.Post("/login", Store.Login())

	// Hubs
	router.Post("/hub/add", Store.CreateHub())
	router.Get("/hubs", Store.ShowHubs())
	router.Get("/hub/{id}", Store.ShowOneHub())
	// router.Delete("/hub/{id}", Store.DeleteHub())

	// Restaurants
	router.Post("/restaurant/add", Store.CreateRestaurant())
	router.Get("/restaurants/{hub_id}", Store.ShowRestaurants())
	router.Get("/restaurant/{id}", Store.ShowOneRestaurant())
	router.Delete("/restaurant/{id}", Store.DeleteRestaurant())

	// Order
	router.Post("/order/finish/{id}", Store.FinishOrder())
	router.Post("/order/status/{id}", Store.StatusUpdate())
	router.Get("/order/{restaurant_id}", Store.ShowOrders())
	router.Get("/order/{state}/{restaurant_id}", Store.ShowStateOrders())
	router.Get("/order/retrieve_code/{state}/{retrieve_code}", Store.ShowOrdersByRetrieveCoce())
	router.Post("/order/add", Store.CreateOrder())

	http.ListenAndServe(":5686", router)
}
