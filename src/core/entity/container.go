package entity

import (
	"time"
)

type ContainerID string

type Container struct {
	ID         ContainerID
	Cmd        []string
	Entrypoint []string
	Created    time.Time
	Name       string

	Image    *Image
}
