package Entity

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"pokomand-go/Middleware"
	"strconv"
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

func AddUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		item := User{}
		// add json data
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			log.Println("Erreur lors de la lecture du corps de la requête:", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// open db
		db := Middleware.OpenDB()

		// call in db
		_, errdb := db.Exec("INSERT INTO Users (last_name,first_name,username,password) VALUES (?,?,?,?)", item.LastName, item.FirstName, item.Username, item.Password)

		if errdb != nil {
			log.Fatal("err3 ", errdb)
		}

		// json response
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Great Success",
		})
	}
}

func GetUserById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Open db
		db := Middleware.OpenDB()

		// select the id in /user/{id}
		queryId := chi.URLParam(request, "id")
		// convert string id in int
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal("err3", err)
			return
		}
		// call in db
		rows, _ := db.Query("SELECT * FROM Users WHERE id = ?", id)
		defer rows.Close()
		users := []User{}

		// select all the rows for user_id
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
