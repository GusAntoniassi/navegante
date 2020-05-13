package dockergateway

import (
	"fmt"
	"github.com/gusantoniassi/navegante/core/entity"
)

func (g *Gateway) ContainerRun(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerStart(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerStop(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerKill(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRestart(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRemove(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRefresh(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}
