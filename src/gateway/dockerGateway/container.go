package dockerGateway

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gusantoniassi/shipmate/core/entity"
)

func hydrateFromTypeContainer(c types.Container) *entity.Container {
	var ec entity.Container

	// For `docker ps` the Command comes in space-separated format, like so: `Entrypoint Cmd1 Cmd2 Cmd3 ..."
	command := c.Command
	commandSplit := strings.Split(command, " ")

	ec.Entrypoint = []string{commandSplit[0]} // Cast the first position to an array
	ec.Cmd = commandSplit[1:]                 // Get everything but the first position
	ec.Created = time.Unix(c.Created, 0)
	ec.Id = entity.ContainerID(c.ID)
	ec.Name = strings.TrimLeft(c.Names[0], "/") // Names come prefixed with their parent, and "/" is the local Docker Daemon

	return &ec
}

func (g *Gateway) ContainerGetAll() ([]*entity.Container, error) {
	ctx := context.Background()

	containers, err := g.Docker.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return nil, err
	}

	var ecs []*entity.Container

	for _, c := range containers {
		ec := hydrateFromTypeContainer(c)
		ecs = append(ecs, ec)
	}

	return ecs, nil
}

func (g *Gateway) ContainerGet(cid entity.ContainerID) (*entity.Container, error) {
	ctx := context.Background()

	filters := filters.NewArgs()
	filters.Add("id", string(cid))

	containers, err := g.Docker.ContainerList(ctx, types.ContainerListOptions{
		Limit:   1,
		Filters: filters,
	})

	if err != nil {
		return nil, err
	}

	if len(containers) <= 0 {
		return nil, nil
	}

	return hydrateFromTypeContainer(containers[0]), nil
}

func (g *Gateway) ContainerRun(c *entity.Container) error {
	return nil
}

func (g *Gateway) ContainerStop(c *entity.Container) error {
	return nil
}

func (g *Gateway) ContainerKill(c *entity.Container) error {
	return nil
}

func (g *Gateway) ContainerRestart(c *entity.Container) error {
	return nil
}

func (g *Gateway) ContainerRemove(c *entity.Container) error {
	return nil
}

func (g *Gateway) ContainerRefresh(c *entity.Container) error {
	return nil
}
