package Entity

import (
	"encoding/json"
	"log"
	"net/http"
	"pokomand-go/Middleware"
)

type Hub struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RestaurantId string `json:"restaurant_id"`
	UserId       string `json:"user_id"`
}

func GetAllHubs() http.HandlerFunc {
	// http à retirer et à mettre dans l'appelle dans Store prendre exemple sur User

	return func(writer http.ResponseWriter, request *http.Request) {
		db := Middleware.OpenDB()
		writer.Header().Set("Content-Type", "application/json")
		rows, _ := db.Query("SELECT * FROM Hubs")
		defer rows.Close()
		hubs := []Hub{}
		for rows.Next() {
			hub := Hub{}
			_ = rows.Scan(&hub.ID, &hub.Name, &hub.RestaurantId, &hub.UserId)
			hubs = append(hubs, hub)
		}

		err := json.NewEncoder(writer).Encode(hubs)
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}
}

// func AddHub() http.HandlerFunc {
// 	// open db
// 	db := Middleware.OpenDB()

// 	// call in db
// 	result, errdb := db.Exec("INSERT INTO Hubs (name,restaurant_id,uder_id) VALUES (?,?,?)", item.Name, item.RestaurantId, item.UserId)

// 	if errdb != nil {
// 		log.Fatal("err3 ", errdb)
// 	}

// 	lastHub, _ := result.LastInsertId()
// 	// json response
// 	return lastHub
// }

type HubsInterface interface {
	GetAllHubs() ([]Hub, error)
	AddHub() ([]Hub, error)
}
