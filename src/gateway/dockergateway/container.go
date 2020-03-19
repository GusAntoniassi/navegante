package dockergateway

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gusantoniassi/navegante/core/entity"
	"strings"
	"time"
)

func hydrateImageFromTypeContainer(c *types.Container) *entity.Image {
	// By default, images come in the format "author/name", without a tag if it's the latest one
	imageName := c.Image
	imageTag := "latest"

	// If the image name contains a ":", the tag is not "latest"
	if strings.Contains(imageName, ":") {
		imageSplit := strings.Split(imageName, ":")
		imageName = imageSplit[0]
		imageTag = imageSplit[1]
	}

	return &entity.Image{
		ID:   c.ImageID,
		Name: imageName,
		Tag:  imageTag,
	}
}

func hydratePortsFromTypePort(p []types.Port) *[]entity.PortMapping {
	ports := make([]entity.PortMapping, 0, len(p))

	for _, port := range p {
		ports = append(ports, entity.PortMapping{
			IP:            port.IP,
			ContainerPort: port.PrivatePort,
			HostPort:      port.PublicPort,
			Protocol:      port.Type,
		})
	}

	return &ports
}

func hydrateVolumesFromTypeMountPoint(mountPoints []types.MountPoint) *[]entity.Volume {
	volumes := make([]entity.Volume, 0, len(mountPoints))

	for _, m := range mountPoints {
		volumes = append(volumes, entity.Volume{
			Name:          m.Name,
			Type:          string(m.Type),
			HostPath:      m.Source,
			ContainerPath: m.Destination,
			Mode:          m.Mode,
			ReadWrite:     m.RW,
		})
	}

	return &volumes
}

func hydrateNetworkFromTypeNetworkSettings(ns types.SummaryNetworkSettings) *[]entity.Network {
	networks := make([]entity.Network, 0, len(ns.Networks))

	for k, n := range ns.Networks {
		networks = append(networks, entity.Network{
			ID:        n.NetworkID,
			Name:      k,
			Gateway:   n.Gateway,
			IPAddress: n.IPAddress,
			Links:     n.Links,
			Aliases:   n.Aliases,
		})
	}

	return &networks
}

func hydrateFromTypeContainer(c *types.Container) *entity.Container {
	var ec entity.Container

	// For `docker ps` the Command comes in space-separated format, like so: `Entrypoint Cmd1 Cmd2 Cmd3 ..."
	command := c.Command
	commandSplit := strings.Split(command, " ")

	ec.Entrypoint = []string{commandSplit[0]} // Cast the first position to an array
	ec.Cmd = commandSplit[1:]                 // Get everything but the first position
	ec.Created = time.Unix(c.Created, 0)
	ec.ID = entity.ContainerID(c.ID)
	ec.Name = strings.TrimLeft(c.Names[0], "/") // Names come prefixed with their parent, and "/" is the local Docker Daemon
	ec.Image = hydrateImageFromTypeContainer(c)
	ec.Ports = hydratePortsFromTypePort(c.Ports)
	ec.Labels = c.Labels
	ec.Volumes = hydrateVolumesFromTypeMountPoint(c.Mounts)
	ec.Networks = hydrateNetworkFromTypeNetworkSettings(*c.NetworkSettings)

	return &ec
}

func (g *Gateway) ContainerGetAll() ([]*entity.Container, error) {
	ctx := context.Background()

	containers, err := g.Docker.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return nil, err
	}

	ecs := make([]*entity.Container, 0, len(containers))

	for _, c := range containers {
		ec := hydrateFromTypeContainer(&c)
		ecs = append(ecs, ec)
	}

	return ecs, nil
}

func (g *Gateway) ContainerGet(cid entity.ContainerID) (*entity.Container, error) {
	ctx := context.Background()

	f := filters.NewArgs()
	f.Add("id", string(cid))

	containers, err := g.Docker.ContainerList(ctx, types.ContainerListOptions{
		Limit:   1,
		Filters: f,
	})

	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, nil
	}

	return hydrateFromTypeContainer(&containers[0]), nil
}

func (g *Gateway) ContainerRun(i *entity.Image) (*entity.Container, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerStart(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerStop(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerKill(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRestart(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRemove(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}

func (g *Gateway) ContainerRefresh(c *entity.Container) error {
	return fmt.Errorf("not implemented yet")
}
