package entity

import (
	"time"
)

type ContainerID string

type PortMapping struct {
	IP            string
	ContainerPort uint16
	HostPort      uint16
	Protocol      string
}

type Container struct {
	ID         ContainerID
	Cmd        []string
	Entrypoint []string
	Created    time.Time
	Name       string

	Image    *Image
	Ports    *[]PortMapping
	Labels   map[string]string
	Volumes  *[]Volume
}
