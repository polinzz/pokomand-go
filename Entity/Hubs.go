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
	HubId        string `json:"hub_id"`
}

func GetHubs() http.HandlerFunc {
	// http à retirer et à mettre dans l'appelle dans Store prendre exemple sur User

	return func(writer http.ResponseWriter, request *http.Request) {
		db := Middleware.OpenDB()
		writer.Header().Set("Content-Type", "application/json")
		rows, _ := db.Query("SELECT * FROM Hubs")
		defer rows.Close()
		hubs := []Hub{}
		for rows.Next() {
			hub := Hub{}
			_ = rows.Scan(&hub.ID, &hub.Name, &hub.RestaurantId, &hub.HubId)
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

type HubsInterface interface {
	GetHubs() ([]Hub, error)
	AddHub() ([]Hub, error)
}
