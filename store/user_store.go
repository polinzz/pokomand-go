package Store

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"pokomand-go/Entity"
)

func Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryUsers := Entity.User{}
		err := json.NewDecoder(request.Body).Decode(&queryUsers)
		if err != nil {
			log.Fatal(err)
		}
		user := Entity.GetUserByUsername(queryUsers.Username)

		if user.Password == queryUsers.Password {
			var store = sessions.NewCookieStore([]byte("your-secret-key"))
			session, _ := store.Get(request, "session-name")

			// Stockez une valeur dans la session
			session.Values["variable_key"] = "valeur"
			session.Save(request, writer)

			err = json.NewEncoder(writer).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "success",
				Message: "Great Success",
			})
		} else {
			err = json.NewEncoder(writer).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "error",
				Message: "error",
			})
		}
	}
}
