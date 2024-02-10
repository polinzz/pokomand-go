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
		log.Println("Début du traitement de la requête CreateHub")
		queryHub := Entity.Hub{}
		log.Println("queryHub:", queryHub)

		err := json.NewDecoder(request.Body).Decode(&queryHub)
		if err != nil {
			log.Fatal(err)
			log.Println("Erreur lors du décodage JSON:", err)
		}

		store := sessions.NewCookieStore([]byte("poko"))
		session, _ := store.Get(request, "session-name")
		log.Println("store:", store)
		log.Println("session:", session)

		// Stockez une valeur dans la session
		userId := session.Values["user_id"].(int64)
		log.Println("userId:", userId)

		lastId := Entity.AddHub(queryHub, userId)
		log.Println("lastId:", lastId)

		hub := Entity.GetHubById(lastId)

		log.Println("hub:", hub)

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

func DeleteHub() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(queryId)
		hub := Entity.DeleteHubByID(int64(id))

		json.NewEncoder(writer).Encode(struct {
			Status  string     `json:"status"`
			Message Entity.Hub `json:"message"`
		}{
			Status:  "success",
			Message: hub,
		})
	}
}
