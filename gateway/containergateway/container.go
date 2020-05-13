package containergateway

import "github.com/gusantoniassi/navegante/core/entity"

type Container interface {
	ContainerGetAll() ([]*entity.Container, error)
	ContainerGet(cid entity.ContainerID) (*entity.Container, error)
	ContainerStatsAll() ([]*entity.Stat, error)
	ContainerStats(cid string) (*entity.Stat, error)
	ContainerRun(c *entity.Container) error
	ContainerStop(c *entity.Container) error
	ContainerKill(c *entity.Container) error
	ContainerRestart(c *entity.Container) error
	ContainerRemove(c *entity.Container) error
	ContainerRefresh(c *entity.Container) error
}
