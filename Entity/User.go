package Entity

import (
	"encoding/json"
	"fmt"
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
	// http à retirer et à mettre dans l'appelle dans Store
	return func(writer http.ResponseWriter, request *http.Request) {
		// call at the db
		db := Middleware.OpenDB()

		// Use the table of the db
		rows, _ := db.Query("SELECT * FROM Users")
		defer rows.Close()

		// initialize User type
		users := []User{}

		// add all the row in users
		for rows.Next() {
			user := User{}
			_ = rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Username, &user.Password, &user.HubId, &user.RestaurantId, &user.Role)
			users = append(users, user)
		}

		// json response
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
	// open db
	db := Middleware.OpenDB()

	hashPassword := Middleware.HashPassword(item.Password)
	// call in db
	result, errdb := db.Exec("INSERT INTO Users (last_name,first_name,username,password) VALUES (?,?,?,?)", item.LastName, item.FirstName, item.Username, hashPassword)

	if errdb != nil {
		log.Fatal("err3 ", errdb)
	}

	lastUser, _ := result.LastInsertId()
	// json response
	return lastUser
}

func GetUserById(id int64) User {
	// Open db
	db := Middleware.OpenDB()
	user := User{}
	// call in db
	err := db.QueryRow("SELECT last_name,first_name,username,role FROM Users WHERE id = ?", id).Scan(&user.LastName, &user.FirstName, &user.Username, &user.Role)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func GetUserByUsername(username string) User {
	// Open db
	fmt.Println(username)
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
