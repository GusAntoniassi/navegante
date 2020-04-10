package container

import (
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
)

type UseCase interface {
	GetAll() ([]*entity.Container, error)
	Get(entity.ContainerID, error)
	Run(c *entity.Container) error
	Stop(c *entity.Container) error
	Kill(c *entity.Container) error
	Restart(c *entity.Container) error
	Remove(c *entity.Container) error
	Refresh(c *entity.Container) error
}

type Service struct {
	containerGw *containergateway.Gateway
}

func NewService(cGw *containergateway.Gateway) *Service {
	return &Service{
		containerGw: cGw,
	}
}

// @TODO: implement concrete usecases
