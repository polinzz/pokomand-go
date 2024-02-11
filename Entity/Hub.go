package Entity

import (
	"log"
	"pokomand-go/Middleware"
)

type Hub struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RestaurantId int    `json:"restaurant_id"`
	UserId       int    `json:"user_id"`
}

func AddHub(item Hub, userId int64) int64 {
	log.Println("Début du traitement de la requête AddHub")

	db := Middleware.OpenDB()

	result, errdb := db.Exec(
		"INSERT INTO Hubs (name,restaurant_id,user_id) VALUES (?,?,?)",
		item.Name, item.RestaurantId, userId,
	)
	log.Println("result", result)

	if errdb != nil {
		log.Fatal(errdb)
	}

	lastHub, _ := result.LastInsertId()

	log.Println("lastHub", lastHub)

	return lastHub
}

func GetHubById(id int64) Hub {
	// Open db
	db := Middleware.OpenDB()
	hub := Hub{}
	// call in db
	err := db.QueryRow("SELECT * FROM Hubs WHERE id = ?", id).Scan(&hub.ID, &hub.Name, &hub.RestaurantId, &hub.UserId)
	if err != nil {
		log.Fatal(err)
	}
	return hub
}

func GetAllHubs() []Hub {
	log.Println("Début du traitement de la requête GetAllHubs")

	// call at the db
	db := Middleware.OpenDB()

	// Use the table of the db
	rows, _ := db.Query("SELECT * FROM Hubs")
	defer rows.Close()

	// initialize User type
	hubs := []Hub{}

	// add all the row in users
	for rows.Next() {
		hub := Hub{}
		_ = rows.Scan(&hub.ID, &hub.Name, &hub.RestaurantId, &hub.UserId)
		hubs = append(hubs, hub)
	}

	return hubs
}

func DeleteHubByID(id int64) {
	// Open db
	db := Middleware.OpenDB()
	// call in db
	db.Exec("DELETE FROM Hubs WHERE id = ?", id)

}

// type HubsInterface interface {
// 	GetAllHubs() ([]Hub, error)
// 	AddHub() ([]Hub, error)
// }
