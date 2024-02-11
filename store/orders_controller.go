package Store

import (
	"encoding/json"
	"log"
	"net/http"
	"pokomand-go/Entity"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func CreateOrder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryOrder := Entity.Order{}

		err := json.NewDecoder(request.Body).Decode(&queryOrder)
		if err != nil {
			log.Fatal(err)
		}
		lastId := Entity.AddOrder(queryOrder)

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

		store := sessions.NewCookieStore([]byte("poko"))
		session, err := store.Get(request, "session-name")

		if err != nil {
			log.Fatal(err)
		}
		userId := session.Values["user_id"].(int64)
		order := Entity.OrderFinish(int64(id), userId)

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
		queryStatus := Entity.Order{}
		err := json.NewDecoder(request.Body).Decode(&queryStatus)
		if err != nil {
			log.Fatal(err)
		}
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}
		store := sessions.NewCookieStore([]byte("poko"))
		session, _ := store.Get(request, "session-name")

		userId := session.Values["user_id"].(int64)

		if err != nil {
			log.Fatal(err)
		}

		order := Entity.ChangeStatus(queryStatus, int64(id), userId)

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
		queryId := chi.URLParam(request, "restaurant_id")
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}
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

func ShowOrdersByRetrieveCode() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "retrieve_code")
		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}
		order := Entity.GetOrderByRetrieveCode(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string       `json:"status"`
			Message Entity.Order `json:"message"`
		}{
			Status:  "success",
			Message: order,
		})
	}
}

func ShowStateOrders() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "restaurant_id")
		queryState := chi.URLParam(request, "state")

		id, err := strconv.Atoi(queryId)
		if err != nil {
			log.Fatal(err)
		}
		var finish bool
		switch queryState {
		case "finish":
			finish = true
		case "not_finish":
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
