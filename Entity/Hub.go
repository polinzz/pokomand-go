package Entity

import (
	"log"
	"pokomand-go/Middleware"
)

type Hub struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"user_id"`
}

func AddHub(item Hub, userId int64) int64 {
	db := Middleware.OpenDB()

	result, errdb := db.Exec(
		"INSERT INTO Hubs (name,user_id) VALUES (?,?)",
		item.Name, userId,
	)

	if errdb != nil {
		log.Fatal(errdb)
	}

	lastHub, _ := result.LastInsertId()

	return lastHub
}

func GetHubById(id int64) Hub {
	// Open db
	db := Middleware.OpenDB()
	hub := Hub{}
	// call in db
	err := db.QueryRow("SELECT * FROM Hubs WHERE id = ?", id).Scan(&hub.ID, &hub.Name, &hub.UserId)
	if err != nil {
		log.Fatal(err)
	}
	return hub
}

func GetAllHubs() []Hub {
	db := Middleware.OpenDB()

	rows, _ := db.Query("SELECT * FROM Hubs")
	defer rows.Close()

	hubs := []Hub{}

	for rows.Next() {
		hub := Hub{}
		_ = rows.Scan(&hub.ID, &hub.Name, &hub.UserId)
		hubs = append(hubs, hub)
	}

	return hubs
}

func DeleteHubByID(id int64) {
	db := Middleware.OpenDB()
	db.Exec("DELETE FROM Hubs WHERE id = ?", id)

}
