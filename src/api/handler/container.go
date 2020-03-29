package handler

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
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
	return nil
}

func getContainer(gw containergateway.Gateway) http.Handler {
	return nil
}
