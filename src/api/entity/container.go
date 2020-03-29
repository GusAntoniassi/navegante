package entity

import (
	"github.com/gusantoniassi/navegante/core/entity"
	"time"
)

type PortMapping struct {
	IP            string
	ContainerPort uint16
	HostPort      uint16
	Protocol      string
}

type Container struct {
	ID         entity.ContainerID `json:"id"`
	Cmd        string             `json:"cmd"`
	Entrypoint string             `json:"entrypoint"`
	Created    time.Time          `json:"created"`
	Name       string             `json:"name"`
	State      string             `json:"state"`
	Status     string             `json:"status"`

	Image    string            `json:"image"`
	Ports    []string          `json:"ports"`
	Labels   map[string]string `json:"labels"`
	Volumes  []string          `json:"volumes"`
	Networks []string          `json:"networks"`
}

func NewContainer(c *entity.Container) Container {
	portMappings := make([]string, len(c.Ports))
	volumes := make([]string, len(c.Volumes))
	networks := make([]string, len(c.Networks))

	for i, p := range c.Ports {
		portMappings[i] = p.String()
	}

	for i, v := range c.Volumes {
		volumes[i] = v.String()
	}

	for i, n := range c.Networks {
		networks[i] = n.String()
	}

	newC := Container{
		ID:         c.ID,
		Cmd:        c.Cmd.String(),
		Entrypoint: c.Entrypoint.String(),
		Created:    c.Created,
		Name:       c.Name,
		State:      c.State,
		Status:     c.Status,
		Image:      c.Image.String(),
		Ports:      portMappings,
		Volumes:    volumes,
		Networks:   networks,
	}

	return newC
}
