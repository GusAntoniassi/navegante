package entity

import (
	"fmt"
	"strings"
	"time"
)

type ContainerID string

type PortMapping struct {
	IP            string
	ContainerPort uint16
	HostPort      uint16
	Protocol      string
}

func (p PortMapping) String() string {
	if p.HostPort == 0 {
		return fmt.Sprintf("%d/%s", p.ContainerPort, p.Protocol)
	}

	return fmt.Sprintf("%d:%d/%s", p.HostPort, p.ContainerPort, p.Protocol)
}

type Cmd []string
type Entrypoint []string

func (c Cmd) String() string {
	return strings.Join(c, " ")
}

func (e Entrypoint) String() string {
	return strings.Join(e, " ")
}

type Container struct {
	ID         ContainerID
	Cmd        Cmd
	Entrypoint Entrypoint
	Created    time.Time
	Name       string
	State      string
	Status     string

	Image    *Image
	Ports    []PortMapping
	Labels   map[string]string
	Volumes  []Volume
	Networks []Network
}
