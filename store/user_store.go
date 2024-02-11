package Store

import (
	"encoding/json"
	"log"
	"net/http"
	"pokomand-go/Entity"
	"pokomand-go/Middleware"
	"strings"

	"github.com/gorilla/sessions"
)

func Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryUsers := Entity.User{}
		err := json.NewDecoder(request.Body).Decode(&queryUsers)
		if err != nil {
			log.Fatal(err)
		}
		user := Entity.GetUserByUsername(queryUsers.Username)
		if strings.Compare(user.Password, Middleware.HashPassword(queryUsers.Password)) == 0 {
			store := sessions.NewCookieStore([]byte("poko"))
			session, _ := store.Get(request, "session-name")

			session.Values["user"] = user.ID
			session.Save(request, writer)

			json.NewEncoder(writer).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "success",
				Message: "Great Success",
			})
		} else {
			json.NewEncoder(writer).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "error",
				Message: "error",
			})
		}
	}
}

func SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		item := Entity.User{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		lastId := Entity.AddUser(item)
		store := sessions.NewCookieStore([]byte("poko"))
		session, _ := store.Get(request, "session-name")

		// Stockez une valeur dans la session
		session.Values["user_id"] = lastId
		session.Save(request, writer)

		user := Entity.GetUserById(lastId)

		json.NewEncoder(writer).Encode(struct {
			Status  string      `json:"status"`
			Message Entity.User `json:"message"`
		}{
			Status:  "success",
			Message: user,
		})
	}
}
