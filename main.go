package main

import (
	"fmt"

	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
	"github.com/gusantoniassi/navegante/gateway/dockergateway"

	"github.com/docker/docker/client"
)

func main() {
	// Example Docker client implementation
	// @TODO: remove this method
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	cGw := containergateway.Container(dockergateway.NewGateway(c))

	//data, err := getContainers(cGw)
	data, err := getStats(cGw)

	if err != nil {
		panic(err)
	}

	//marshalled, _ := json.Marshal(data)
	//fmt.Printf("%s\n", marshalled)

	for _, v := range data {
		fmt.Printf(
			"ID: %s\n"+
				"CPU%%: %.2f\n"+
				"Mem%%: %.2f\n"+
				"Mem usg/lim: %d/%d\n"+
				"Net I/O: %d/%d\n"+
				"Block I/O: %d/%d\n",
			v.ContainerID[:12],
			v.CPUPercent,
			v.MemoryPercent,
			v.MemoryUsage,
			v.MemoryTotal,
			v.NetworkInput,
			v.NetworkOutput,
			v.BlockRead,
			v.BlockWrite,
		)
	}
}

func getContainers(cGw containergateway.Container) ([]*entity.Container, error) {
	containers, err := cGw.ContainerGetAll()

	return containers, err
}

func getStats(cGw containergateway.Container) ([]*entity.Stat, error) {
	stats, err := cGw.ContainerStatsAll()

	return stats, err
}
