package handler

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	apiEntity "github.com/gusantoniassi/navegante/api/entity"
	"github.com/gusantoniassi/navegante/core/entity/user"
	"log"
	"net/http"
	"strconv"
)

func MakeUserHandlers(r *mux.Router, n *negroni.Negroni, mgr user.Manager) {
	r.Handle("/users", n.With(
		negroni.Wrap(getAllUsers(mgr)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/users", n.With(
		negroni.Wrap(addUser(mgr)),
	)).Methods("POST")
}

func getAllUsers(mgr user.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := mgr.List()

		if err != nil {
			log.Println("error calling mgr.List: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error getting users from the database"))
			return
		}

		if len(users) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError("No users registered"))
			return
		}

		apiUsers := make([]apiEntity.User, len(users))
		for i, usr := range users {
			apiUsers[i] = apiEntity.User(*usr)
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(apiUsers)

		if err != nil {
			log.Println("error converting users to JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error converting users to JSON"))
			return
		}
	})
}

func addUser(mgr user.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var usr user.User

		err := json.NewDecoder(r.Body).Decode(&usr)

		if err != nil {
			log.Println("error decoding user JSON: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError("Error decoding user, please check the request body and try again"))
			return
		}

		id, err := mgr.Create(&usr)
		if err != nil {
			log.Println("error creating user: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error saving user"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("{\"id\": \"%s\"}", strconv.FormatUint(uint64(id), 10))))
	})
}
