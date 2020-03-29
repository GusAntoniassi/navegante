package entity

import (
	"fmt"
)

type Volume struct {
	Name        string
	Type        string
	Source      string
	Destination string
	Mode        string
	ReadWrite   bool
}

// formats a volume struct into a format similar to the one used in Docker Compose
func (v Volume) String() string {
	if v.Type == "bind" {
		return fmt.Sprintf("%s:%s:%s", v.Source, v.Destination, v.Mode)
	}

	// Named volumes have a source pointing to /var/lib/docker/volumes
	if v.Type == "volume" && v.Source != "" {
		return fmt.Sprintf("%s:%s", v.Name, v.Destination)
	}

	// For volumes without a name, return just the destination path
	return v.Destination
}
