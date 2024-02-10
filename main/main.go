package main

import (
	"net/http"
	"pokomand-go/Entity"
	Store "pokomand-go/Store"

	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := chi.NewRouter()

	// ENDPOINTS

	// Users
	router.Get("/user", Entity.GetAllUsers())
	router.Post("/user/add", Store.SignUp())
	router.Post("/login", Store.Login())

	// Hubs
	router.Post("/hub/add", Store.CreateHub())
	router.Get("/hubs", Store.ShowHubs())
	router.Get("/hub/{id}", Store.ShowOneHub())
	router.Delete("/hub/{id}", Store.DeleteHub())

	// Restaurants
	router.Post("/restaurant/add", Entity.CreateRestaurantHandler)
	router.Get("/restaurants", Entity.GetAllRestaurants)
	router.Delete("/restaurant", Entity.DeleteRestaurantByID)

	// Order
	router.Patch("/order/finish/{id}", Store.FinishOrder())
	router.Patch("/order/status/{id}", Store.StatusUpdate())
	router.Get("/order/{id}", Store.ShowOrders())
	router.Get("/order/{state}/{id}", Store.ShowStateOrders())
	router.Put("/order/add", Store.CreateOrder())

	http.ListenAndServe(":5686", router)
}

// FAKE DATA FOR RESTAURANTS

// Create http://localhost:5686/restaurant/add
// {
//   "name": "Nom du restaurant 1",
//   "foods": [
//     {
//         "name": "food 1",
//         "price" : "2"
//     },
//     {
//         "name": "food 2",
//         "price" : "3"
//     }
//   ],
//   "drinks": [
//     {
//         "name": "drink 1",
//         "price" : "2"
//     },
//     {
//         "name": "drink 2",
//         "price" : "3"
//     }
//   ]
// }

// Delete http://localhost:5686/restaurant?id={id}
