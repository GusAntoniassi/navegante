package entity

import (
	"time"
)

type ContainerID string

type Container struct {
	Id         ContainerID
	Cmd        []string
	Entrypoint []string
	Created    time.Time
	Name       string
}
