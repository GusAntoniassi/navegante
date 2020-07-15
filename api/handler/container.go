package handler

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	apiEntity "github.com/gusantoniassi/navegante/api/entity"
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
	"log"
	"net/http"
)

func MakeContainerHandlers(r *mux.Router, n *negroni.Negroni, gw containergateway.Gateway) {
	r.Handle("/containers", n.With(
		negroni.Wrap(getAllContainers(gw)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/containers/{id}", n.With(
		negroni.Wrap(getContainer(gw)),
	)).Methods("GET", "OPTIONS")
}

func getAllContainers(gw containergateway.Gateway) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gwContainers, err := gw.ContainerGetAll()
		returnStats := r.FormValue("stats") == "true"

		if err != nil {
			log.Println("error calling gw.ContainerGetAll: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error getting the containers from the API"))
			return
		}

		if len(gwContainers) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError("No containers running"))
			return
		}

		containers := make([]apiEntity.Container, len(gwContainers))
		statResults := make([]chan *entity.Stat, len(gwContainers))

		for i, c := range gwContainers {
			containers[i] = apiEntity.NewContainer(c)

			if returnStats {
				ch := make(chan *entity.Stat)
				statResults[i] = ch
				go getStats(gw, string(c.ID), ch)
			}
		}

		if returnStats {
			for i := range statResults {
				for stats := range statResults[i] {
					containers[i].Statistics = stats
				}
			}
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(containers)

		if err != nil {
			log.Println("error converting container to JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error converting container to JSON"))
			return
		}
	})
}

func getContainer(gw containergateway.Gateway) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		returnStats := r.FormValue("stats") == "true"

		gwContainer, err := gw.ContainerGet(entity.ContainerID(id))

		if err != nil {
			log.Println("error calling gw.ContainerGetAll: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error getting the containers from the API"))
			return
		}

		if gwContainer == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError(fmt.Sprintf("Container '%s' not found", id)))
			return
		}

		container := apiEntity.NewContainer(gwContainer)

		if returnStats {
			stats := getStatsSync(gw, id)
			container.Statistics = stats
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(container)

		if err != nil {
			log.Println("error converting container to JSON: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Error converting container to JSON"))
			return
		}
	})
}

func getStatsSync(gw containergateway.Gateway, id string) *entity.Stat {
	stats, err := gw.ContainerStats(id)

	if err != nil {
		log.Println("error calling gw.ContainerGetStats: ", err)
		return nil
	}

	return stats
}

func getStats(gw containergateway.Gateway, id string, c chan *entity.Stat) {
	defer close(c)
	stats := getStatsSync(gw, id)
	c <- stats
}
