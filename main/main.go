package main

import (
	"net/http"
	"pokomand-go/Entity"
	"pokomand-go/store"

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
	router.Get("/hubs", Entity.GetAllHubs())

	// Restaurants
	router.Post("/restaurant/add", Entity.CreateRestaurantHandler)
	router.Get("/restaurants", Entity.GetAllRestaurants)
	router.Delete("/restaurant", Entity.DeleteRestaurantByID)

	// Order
	router.Post("/order/finish/{id}", Store.FinishOrder())
	router.Post("/order/status/{id}", Store.StatusUpdate())
	router.Get("/order/{restaurant_id}", Store.ShowOrders())
	router.Get("/order/{state}/{restaurant_id}", Store.ShowStateOrders())
	router.Get("/order/retrieve_code/{state}/{retrieve_code}", Store.ShowOrdersByRetrieveCoce())
	router.Post("/order/add", Store.CreateOrder())

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
