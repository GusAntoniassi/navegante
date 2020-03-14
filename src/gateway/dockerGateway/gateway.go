package dockerGateway

import "github.com/docker/docker/client"

type Gateway struct {
	Docker client.CommonAPIClient
}

func NewGateway(docker client.CommonAPIClient) *Gateway {
	return &Gateway{
		Docker: docker,
	}
}
