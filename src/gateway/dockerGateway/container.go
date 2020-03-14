package dockerGateway

import (
	"github.com/gusantoniassi/shipmate/core/entity"
)

func (g *Gateway) ContainerGetAll() ([]*entity.Container, error) {
	return nil, nil
}

func (g *Gateway) ContainerGet(cid entity.ContainerID) (*entity.Container, error) {
	return nil, nil
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
