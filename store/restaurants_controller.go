package Store

import (
	"encoding/json"
	"log"
	"net/http"
	"pokomand-go/Entity"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryRestaurant := Entity.Restaurant{}

		err := json.NewDecoder(request.Body).Decode(&queryRestaurant)
		if err != nil {
			log.Fatal(err)
		}
		lastId := Entity.AddRestaurant(queryRestaurant)

		restaurant := Entity.GetRestaurantById(lastId)

		json.NewEncoder(writer).Encode(struct {
			Status  string            `json:"status"`
			Message Entity.Restaurant `json:"message"`
		}{
			Status:  "success",
			Message: restaurant,
		})
	}
}

func ShowOneRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(queryId)
		restaurant := Entity.GetRestaurantById(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string            `json:"status"`
			Message Entity.Restaurant `json:"message"`
		}{
			Status:  "success",
			Message: restaurant,
		})
	}
}

func ShowRestaurants() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "hub_id")
		id, _ := strconv.Atoi(queryId)
		restaurants := Entity.GetAllRestaurantsByHubId(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string              `json:"status"`
			Message []Entity.Restaurant `json:"message"`
		}{
			Status:  "success",
			Message: restaurants,
		})
	}
}

func DeleteRestaurant() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(queryId)
		Entity.DeleteRestaurantByID(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Restaurant supprimé",
		})
	}
}
