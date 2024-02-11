package Entity

import (
	"encoding/json"
	"log"
	"pokomand-go/Middleware"
)

type Restaurant struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	HubId  int     `json:"hub_id"`
	Foods  []Food  `json:"foods"`
	Drinks []Drink `json:"drinks"`
}

type Food struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Drink struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func AddRestaurant(item Restaurant) int64 {
	log.Println("Début du traitement de la requête AddRestaurant")

	db := Middleware.OpenDB()

	foods, err := json.Marshal(&item.Foods)
	drinks, err := json.Marshal(&item.Drinks)
	if err != nil {
		log.Fatal(err)
	}

	result, errdb := db.Exec(
		"INSERT INTO Restaurants (name, hub_id, foods, drinks) VALUES (?, ?, ?, ?)",
		item.Name, item.HubId, foods, drinks,
	)
	log.Println("result", result)

	if errdb != nil {
		log.Fatal(errdb)
	}

	lastRestaurant, _ := result.LastInsertId()

	log.Println("lastRestaurant", lastRestaurant)

	return lastRestaurant
}

func GetRestaurantById(id int64) Restaurant {
	db := Middleware.OpenDB()
	restaurant := Restaurant{}
	var foodsJSON string
	var drinksJSON string

	err := db.QueryRow("SELECT * FROM Restaurants WHERE id = ?", id).
		Scan(&restaurant.ID, &restaurant.Name, &restaurant.HubId, &foodsJSON, &drinksJSON)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(foodsJSON), &restaurant.Foods)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(drinksJSON), &restaurant.Drinks)
	if err != nil {
		log.Fatal(err)
	}

	return restaurant
}

func GetAllRestaurantsByHubId(hubId int64) []Restaurant {
	db := Middleware.OpenDB()
	log.Println("hubId", hubId)

	rows, _ := db.Query("SELECT * FROM Restaurants WHERE hub_id = ?", hubId)
	defer rows.Close()

	log.Println("rows", rows)

	var foodsJSON string
	var drinksJSON string
	restaurants := []Restaurant{}

	for rows.Next() {
		restaurant := Restaurant{}
		_ = rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.HubId, &foodsJSON, &drinksJSON)
		errFoods := json.Unmarshal([]byte(foodsJSON), &restaurant.Foods)
		errDrinks := json.Unmarshal([]byte(drinksJSON), &restaurant.Drinks)
		if errFoods != nil {
			log.Fatal("errFoods ", errFoods)
		}
		if errDrinks != nil {
			log.Fatal("errDrinks ", errDrinks)
		}
		restaurants = append(restaurants, restaurant)
		log.Println("restaurants", restaurants)
	}

	return restaurants
}

func DeleteRestaurantByID(id int64) {
	db := Middleware.OpenDB()
	db.Exec("DELETE FROM Restaurants WHERE id = ?", id)
}
