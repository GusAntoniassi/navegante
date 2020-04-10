package main

import (
	"encoding/json"
	"fmt"
	"github.com/gusantoniassi/navegante/gateway/containergateway"

	"github.com/docker/docker/client"
	"github.com/gusantoniassi/navegante/gateway/dockergateway"
)

func main() {
	// Example Docker client implementation
	// @TODO: remove this

	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	var cGw containergateway.Container

	cGw = dockergateway.NewGateway(c)
	containers, err := cGw.ContainerGetAll()

	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(containers)
	fmt.Printf("%s\n", data)
}
