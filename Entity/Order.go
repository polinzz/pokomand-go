package Entity

import (
	"encoding/json"
	"log"
	"math/rand"
	"pokomand-go/Middleware"
)

type Order struct {
	Id           int        `json:"id"`
	Product      []Products `json:"products"`
	Price        float32    `json:"price"`
	Status       string     `json:"status"`
	IsFinish     bool       `json:"is_finish"`
	RestaurantId int        `json:"restaurant_id"`
	UserId       int        `json:"user_id"`
	RetrieveCode int        `json:"retrieve_code"`
}

type Products struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func AddOrder(item Order, userId int64) int64 {
	db := Middleware.OpenDB()

	products, err := json.Marshal(&item.Product)
	if err != nil {
		log.Fatal(err)
	}
	RetrieveCode := rand.Intn(9000) + 1000
	result, errdb := db.Exec("INSERT INTO Orders (product,restaurant_id,price,status,is_finish,user_id,retrieve_code) VALUES (?,?,?,?,?,?,?)",
		products, item.RestaurantId, item.Price, "En attente", item.IsFinish, userId, RetrieveCode)

	if errdb != nil {
		log.Fatal(errdb)
	}

	lastOrder, _ := result.LastInsertId()

	return lastOrder
}

func GetOrderById(id int64) Order {
	db := Middleware.OpenDB()
	order := Order{}
	var productJSON string
	err := db.QueryRow("SELECT id, product, restaurant_id, price, status, is_finish, user_id, retrieve_code FROM Orders WHERE id = ?", id).Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.UserId, &order.RetrieveCode)

	if err != nil {
		log.Fatal(err)
	}

	// Utiliser json.Unmarshal pour déserialiser la chaîne JSON dans le champ Product
	err = json.Unmarshal([]byte(productJSON), &order.Product)
	if err != nil {
		log.Fatal(err)
	}

	return order
}

func OrderFinish(id int64) Order {
	db := Middleware.OpenDB()

	_, err := db.Exec("UPDATE Orders SET is_finish = true WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	order := GetOrderById(id)

	return order
}

func ChangeStatus(status string, id int64) Order {
	db := Middleware.OpenDB()

	_, err := db.Exec("UPDATE Orders SET status = ? WHERE id = ?", status, id)
	if err != nil {
		log.Fatal(err)
	}
	order := GetOrderById(id)

	return order
}

func GetAllOrders(id int64) []Order {
	// call at the db
	db := Middleware.OpenDB()

	// Use the table of the db
	rows, _ := db.Query("SELECT * FROM Orders WHERE restaurant_id = ?", id)
	defer rows.Close()

	var productJSON string
	// initialize User type
	orders := []Order{}

	// add all the row in users
	for rows.Next() {
		order := Order{}
		_ = rows.Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.UserId, &order.RetrieveCode)
		err := json.Unmarshal([]byte(productJSON), &order.Product)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	return orders

}

func GetOrders(id int64, finish bool) []Order {
	db := Middleware.OpenDB()

	// Use the table of the db
	rows, _ := db.Query("SELECT * FROM Orders WHERE restaurant_id = ? AND is_finish = ?", id, finish)
	defer rows.Close()

	var productJSON string
	// initialize User type
	orders := []Order{}

	// add all the row in users
	for rows.Next() {
		order := Order{}
		_ = rows.Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.UserId, &order.RetrieveCode)
		err := json.Unmarshal([]byte(productJSON), &order.Product)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	return orders
}
