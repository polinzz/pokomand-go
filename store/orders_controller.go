package Store

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"pokomand-go/Entity"
	"strconv"
)

func CreateOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryOrder := Entity.Order{}

		err := json.NewDecoder(request.Body).Decode(&queryOrder)
		if err != nil {
			log.Fatal(err)
		}

		store := sessions.NewCookieStore([]byte("poko"))
		session, _ := store.Get(request, "session-name")

		// Stockez une valeur dans la session
		userId := session.Values["user_id"].(int64)

		lastId := Entity.AddOrder(queryOrder, userId)

		order := Entity.GetOrderById(lastId)

		json.NewEncoder(writer).Encode(struct {
			Status  string       `json:"status"`
			Message Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: order,
		})
	}
}

func FinishOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}

		order := Entity.OrderFinish(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string       `json:"status"`
			Message Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: order,
		})
	}
}

func StatusUpdate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		var queryStatus string
		err := json.NewDecoder(request.Body).Decode(&queryStatus)
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}

		order := Entity.ChangeStatus(queryStatus, int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string       `json:"status"`
			Message Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: order,
		})
	}
}

func ShowOrders() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(queryId)
		orders := Entity.GetAllOrders(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string         `json:"status"`
			Message []Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: orders,
		})
	}
}

func ShowStateOrders() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		queryState := chi.URLParam(request, "state")

		id, _ := strconv.Atoi(queryId)
		var finish bool
		switch queryState {
		case "active":
			finish = true
		case "deactive":
			finish = false
		default:
			log.Fatal("finish mal défini")

		}

		orders := Entity.GetOrders(int64(id), finish)

		json.NewEncoder(writer).Encode(struct {
			Status  string         `json:"status"`
			Message []Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: orders,
		})
	}

}
