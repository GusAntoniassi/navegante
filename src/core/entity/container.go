package entity

import (
	"fmt"
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

type Container struct {
	ID         ContainerID
	Cmd        []string
	Entrypoint []string
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
