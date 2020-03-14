package main

import (
	"fmt"

	"github.com/docker/docker/client"
	"github.com/gusantoniassi/navegante/gateway/dockergateway"
)

func main() {
	// Example Docker client implementation
	// @TODO: remove this

	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	cGw := dockergateway.NewGateway(client)
	containers, err := cGw.ContainerGetAll()

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%+v", container)
	}
}
