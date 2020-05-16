package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
)

func MakeStatHandlers(r *mux.Router, n *negroni.Negroni, gw containergateway.Gateway) {
	r.Handle("/stats", n.With(
		negroni.Wrap(getAllStats(gw)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/stats/{id}", n.With(
		negroni.Wrap(getContainerStats(gw)),
	)).Methods("GET", "OPTIONS")
}

func getAllStats(gw containergateway.Gateway) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stats, err := gw.ContainerStatsAll()

		if err != nil {
			log.Println("error calling gw.ContainerStatsAll: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error getting the containers stats from the API"))
			return
		}

		if len(stats) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError("No containers running"))
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(stats)

		if err != nil {
			log.Println("error converting stats to JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error converting stats to JSON"))
			return
		}
	})
}

func getContainerStats(gw containergateway.Gateway) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		stats, err := gw.ContainerGet(entity.ContainerID(id))

		if err != nil {
			log.Print("error calling gw.ContainerGet: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error getting the stats from the API"))
			return
		}

		if stats == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError(fmt.Sprintf("Container '%s' not found", id)))
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(stats)

		if err != nil {
			log.Print("error converting stats to JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error converting stats to JSON"))
			return
		}
	})
}
