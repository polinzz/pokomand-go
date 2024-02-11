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
	RetrieveCode int        `json:"retrieve_code"`
}

type Products struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

func AddOrder(item Order) int64 {
	db := Middleware.OpenDB()
	totalPrice := 0.0
	for _, value := range item.Product {
		totalPrice = totalPrice + (value.Price * value.Quantity)
	}

	products, err := json.Marshal(&item.Product)

	if err != nil {
		log.Fatal(err)
	}
	RetrieveCode := rand.Intn(9000) + 1000
	result, errdb := db.Exec("INSERT INTO Orders (product,restaurant_id,price,status,retrieve_code) VALUES (?,?,?,?,?)",
		products, item.RestaurantId, totalPrice, "En attente", RetrieveCode)

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
	err := db.QueryRow("SELECT id, product, restaurant_id, price, status, is_finish, retrieve_code FROM Orders WHERE id = ?", id).
		Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.RetrieveCode)

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

func OrderFinish(id int64, userId int64) Order {
	db := Middleware.OpenDB()
	user := GetUserById(userId)

	order := GetOrderById(id)

	if user.Role != "costumers" && user.RestaurantId == order.RestaurantId {
		_, err := db.Exec("UPDATE Orders SET is_finish = true WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("User pas bon")
	}

	UpdateOrder := GetOrderById(id)
	return UpdateOrder
}

func ChangeStatus(queryOrder Order, id int64, userId int64) Order {
	db := Middleware.OpenDB()
	user := GetUserById(userId)
	order := GetOrderById(id)

	if user.Role != "costumers" && user.RestaurantId == order.RestaurantId {
		_, err := db.Exec("UPDATE Orders SET status = ? WHERE id = ?", queryOrder.Status, id)
		if err != nil {
			log.Fatal(err)
		}
	}

	UpdateOrder := GetOrderById(id)
	return UpdateOrder
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
		_ = rows.Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.RetrieveCode)
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
		_ = rows.Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.RetrieveCode)
		err := json.Unmarshal([]byte(productJSON), &order.Product)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	return orders
}

func GetOrderByRetrieveCode(retrieveCode int64) Order {
	db := Middleware.OpenDB()

	order := Order{}
	var productJSON string
	err := db.QueryRow("SELECT id, product, restaurant_id, price, status, is_finish, retrieve_code FROM Orders WHERE retrieve_code = ?", retrieveCode).Scan(&order.Id, &productJSON, &order.RestaurantId, &order.Price, &order.Status, &order.IsFinish, &order.RetrieveCode)

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
