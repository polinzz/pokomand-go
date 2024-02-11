package Entity

import (
	"encoding/json"
	"log"
	"net/http"
	"pokomand-go/Middleware"
)

type User struct {
	ID           int    `json:"id"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	HubId        int    `json:"hub_id"`
	RestaurantId int    `json:"restaurant_id"`
	Role         string `json:"role"`
}

func GetAllUsers() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := Middleware.OpenDB()

		rows, _ := db.Query("SELECT * FROM Users")
		defer rows.Close()

		users := []User{}

		for rows.Next() {
			user := User{}
			_ = rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Username, &user.Password, &user.HubId, &user.RestaurantId, &user.Role)
			users = append(users, user)
		}

		writer.Header().Set("Content-Type", "application/json")
		errUser := json.NewEncoder(writer).Encode(users)
		if errUser != nil {
			log.Fatal(errUser)
			return
		}
		return
	}
}

func AddUser(item User) int64 {
	db := Middleware.OpenDB()

	hashPassword := Middleware.HashPassword(item.Password)
	result, errdb := db.Exec("INSERT INTO Users (last_name,first_name,username,password,restaurant_id) VALUES (?,?,?,?,?)", item.LastName, item.FirstName, item.Username, hashPassword, item.RestaurantId)

	if errdb != nil {
		log.Fatal("err3 ", errdb)
	}

	lastUser, _ := result.LastInsertId()
	return lastUser
}

func GetUserById(id int64) User {
	db := Middleware.OpenDB()
	user := User{}
	err := db.QueryRow("SELECT last_name,first_name,username,role,restaurant_id FROM Users WHERE id = ?", id).Scan(&user.LastName, &user.FirstName, &user.Username, &user.Role, &user.RestaurantId)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func GetUserByUsername(username string) User {
	db := Middleware.OpenDB()
	user := User{}
	err := db.QueryRow("SELECT first_name,last_name,username,password FROM Users WHERE username = ?", username).Scan(&user.LastName, &user.FirstName, &user.Username, &user.Password)
	if err != nil {
		log.Fatal(user)
	}

	return user
}

type UserInterface interface {
	GetAllUsers() ([]User, error)
	AddUser() ([]User, error)
	ShowUser() ([]User, error)
}
