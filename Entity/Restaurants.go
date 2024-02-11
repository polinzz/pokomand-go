package Entity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokomand-go/Middleware"
	"strconv"
)

type Restaurant struct {
	ID     int    `json:"id"`
	Name   string  `json:"name"`
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

func CreateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	var restaurant Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}

	db := Middleware.OpenDB()
	defer db.Close()

	foodsJSON, err := json.Marshal(restaurant.Foods)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	drinksJSON, err := json.Marshal(restaurant.Drinks)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	_, err = db.Exec("INSERT INTO Restaurants (name, foods, drinks) VALUES (?, ?, ?)",
			restaurant.Name, foodsJSON, drinksJSON)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	fmt.Fprintf(w, "Restaurant créé avec succès")
}

func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	db := Middleware.OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, foods, drinks FROM Restaurants")
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer rows.Close()

	var restaurants []Restaurant
	for rows.Next() {
			var restaurant Restaurant
			var foodsJSON, drinksJSON string
			err := rows.Scan(&restaurant.ID, &restaurant.Name, &foodsJSON, &drinksJSON)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

			err = json.Unmarshal([]byte(foodsJSON), &restaurant.Foods)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

			err = json.Unmarshal([]byte(drinksJSON), &restaurant.Drinks)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

			restaurants = append(restaurants, restaurant)
	}
	if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

func DeleteRestaurantByID(w http.ResponseWriter, r *http.Request) {
	restaurantIDStr := r.URL.Query().Get("id")
	restaurantID, err := strconv.Atoi(restaurantIDStr)
	if err != nil {
			http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
			return
	}

	db := Middleware.OpenDB()
	defer db.Close()

	result, err := db.Exec("DELETE FROM Restaurants WHERE id = ?", restaurantID)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	if rowsAffected == 0 {
			http.Error(w, fmt.Sprintf("No restaurant found with ID %d", restaurantID), http.StatusNotFound)
			return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Restaurant with ID %d deleted successfully", restaurantID)
}

func GetRestaurantByID(w http.ResponseWriter, r *http.Request) {
	restaurantIDStr := r.URL.Query().Get("id")
	restaurantID, err := strconv.Atoi(restaurantIDStr)
	if err != nil {
		http.Error(w, "Invalid restaurant ID", http.StatusBadRequest)
		return
	}

	db := Middleware.OpenDB()
	defer db.Close()

	var restaurantName, foodsJSON, drinksJSON string
	err = db.QueryRow("SELECT name, foods, drinks FROM Restaurants WHERE id = ?", restaurantID).Scan(&restaurantName, &foodsJSON, &drinksJSON)
	if err != nil {
		http.Error(w, "Restaurant not found", http.StatusNotFound)
		return
	}

	restaurantInfo := struct {
		Name   string  `json:"name"`
		Foods  []Food  `json:"foods"`
		Drinks []Drink `json:"drinks"`
	}{
		Name:   restaurantName,
		Foods:  make([]Food, 0),
		Drinks: make([]Drink, 0),
	}

	err = json.Unmarshal([]byte(foodsJSON), &restaurantInfo.Foods)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(drinksJSON), &restaurantInfo.Drinks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(restaurantInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
