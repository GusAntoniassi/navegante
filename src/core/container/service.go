package container

import (
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containerGateway"
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
	containerGw *containerGateway.Gateway
}

func NewService(cGw *containerGateway.Gateway) *Service {
	return &Service{
		containerGw: cGw,
	}
}

// @TODO: implement concrete usecases
