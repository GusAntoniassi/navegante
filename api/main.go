package main

import (
	"fmt"
	"github.com/gusantoniassi/navegante/core/entity/user"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"github.com/gusantoniassi/navegante/api/handler"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
	"github.com/gusantoniassi/navegante/gateway/dockergateway"
)

const PORT = 5000

func main() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
		negroni.HandlerFunc(handler.AddContentType),
		negroni.HandlerFunc(handler.AllowCORS),
	)

	gw, err := getDockerGateway()

	if err != nil {
		log.Fatal(err)
	}

	r.NotFoundHandler = http.HandlerFunc(handler.GetNotFoundHandler)

	api1 := r.PathPrefix("/v1").Subrouter()
	handler.MakeContainerHandlers(api1, n, *gw)
	handler.MakeStatHandlers(api1, n, *gw)

	userRepo := user.NewInMemRepository()
	userManager := user.NewManager(userRepo)
	handler.MakeUserHandlers(api1, n, userManager)

	http.Handle("/", r)

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "[error] ", log.Lshortfile)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		ReadTimeout:  handler.REQUEST_TIMEOUT * time.Second,
		WriteTimeout: handler.REQUEST_TIMEOUT * time.Second,
		ErrorLog:     logger,
		Handler:      http.DefaultServeMux,
	}

	fmt.Printf("Starting server listening on port %d\n", PORT)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func getDockerGateway() (*containergateway.Gateway, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return nil, err
	}

	var cGw containergateway.Gateway = dockergateway.NewGateway(c)

	return &cGw, nil
}
