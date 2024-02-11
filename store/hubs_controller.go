package Store

import (
	"encoding/json"
	"log"
	"net/http"
	Entity "pokomand-go/Entity"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func CreateHub() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryHub := Entity.Hub{}

		err := json.NewDecoder(request.Body).Decode(&queryHub)
		if err != nil {
			log.Fatal(err)
		}

		store := sessions.NewCookieStore([]byte("poko"))
		session, _ := store.Get(request, "session-name")

		// Stockez une valeur dans la session
		userId := session.Values["user_id"].(int64)

		lastId := Entity.AddHub(queryHub, userId)

		hub := Entity.GetHubById(lastId)

		json.NewEncoder(writer).Encode(struct {
			Status  string     `json:"status"`
			Message Entity.Hub `json:"message"`
		}{
			Status:  "success",
			Message: hub,
		})
	}
}

func ShowHubs() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hubs := Entity.GetAllHubs()

		json.NewEncoder(writer).Encode(struct {
			Status  string       `json:"status"`
			Message []Entity.Hub `json:"message"`
		}{
			Status:  "success",
			Message: hubs,
		})
	}
}

func ShowOneHub() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(queryId)
		hub := Entity.GetHubById(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string     `json:"status"`
			Message Entity.Hub `json:"message"`
		}{
			Status:  "success",
			Message: hub,
		})
	}
}

// func DeleteHub() http.HandlerFunc {
// 	return func(writer http.ResponseWriter, request *http.Request) {
// 		queryId := chi.URLParam(request, "id")
// 		id, _ := strconv.Atoi(queryId)
// 		hub := Entity.DeleteHubByID(int64(id))

// 		json.NewEncoder(writer).Encode(struct {
// 			Status  string `json:"status"`
// 			Message string `json:"message"`
// 		}{
// 			Status:  "success",
// 			Message: "Hub supprimé",
// 		})
// 	}
// }
