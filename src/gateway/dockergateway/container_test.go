package dockergateway

import (
	"context"
	networktypes "github.com/docker/docker/api/types/network"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/gojuno/minimock/v3"
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/stretchr/testify/assert"
)

func getMockContainers() []types.Container {
	return []types.Container{
		{
			ID:      "0123abcd456e",
			Names:   []string{"/smart_einstein"},
			Image:   "registry.foo.bar/foo:latest",
			ImageID: "sha256:abcde123456",
			Command: "/bin/bash echo foo",
			Created: 1583610097,
			Ports: []types.Port{
				{
					IP:          "0.0.0.0",
					PrivatePort: 80,
					PublicPort:  8080,
					Type:        "tcp",
				},
			},
			SizeRw:     0,
			SizeRootFs: 0,
			Labels: map[string]string{
				"foo": "bar",
			},
			State:  "running",
			Status: "Up 4 hours",
			HostConfig: struct {
				NetworkMode string `json:",omitempty"`
			}{},
			NetworkSettings: &types.SummaryNetworkSettings{
				Networks: map[string]*networktypes.EndpointSettings{
					"bridge": {
						NetworkID: "abcd123456",
						Gateway:   "172.17.0.1",
						IPAddress: "172.17.0.2",
					},
				},
			},
			Mounts: []types.MountPoint{
				{
					Type:        "bind",
					Name:        "",
					Source:      "/foo/bar",
					Destination: "/container/bar",
					Driver:      "",
					Mode:        "rw",
					RW:          true,
					Propagation: "rprivate",
				},
				{
					Type:        "volume",
					Name:        "log",
					Source:      "/var/lib/docker/volumes/foo-log/_data",
					Destination: "/var/log",
					Driver:      "local",
					Mode:        "rw",
					RW:          true,
					Propagation: "",
				},
			},
		},
		{
			ID:      "0321dcba654e",
			Names:   []string{"/silly_bach"},
			Image:   "registry.foo.bar/bar:1.0.2",
			ImageID: "sha256:abcde123123",
			Command: "/bin/bash echo bar",
			Created: 1583610006,
			Ports: []types.Port{
				{PrivatePort: 80, Type: "tcp"},
			},
			SizeRw:     0,
			SizeRootFs: 0,
			Labels:     nil,
			State:      "running",
			Status:     "Up 10 hours",
			HostConfig: struct {
				NetworkMode string `json:",omitempty"`
			}{},
			NetworkSettings: nil,
			Mounts:          nil,
		},
	}
}

func TestGateway_ContainerGetAll(t *testing.T) {
	mockContainers := getMockContainers()

	mc := minimock.NewController(t)
	dockerMock := NewCommonAPIClientMock(mc).ContainerListMock.Set(func(context.Context, types.ContainerListOptions) ([]types.Container, error) {
		return mockContainers, nil
	})

	gw := NewGateway(dockerMock)

	containers, err := gw.ContainerGetAll()

	assert.Nilf(t, err, "ContainerGetAll returns no error")
	assert.NotEmptyf(t, containers, "Should return at least one container")

	for i, container := range containers {
		assert.Equal(t, string(container.ID), mockContainers[i].ID,
			"Container IDs match")
		assert.Equal(t, container.Created, time.Unix(mockContainers[i].Created, 0),
			"Container creation times match")
		assert.Equal(t, container.Name, strings.TrimLeft(mockContainers[i].Names[0], "/"),
			"Container names match")
	}
}

func TestGateway_ContainerGet(t *testing.T) {
	mockContainer := getMockContainers()[0]

	mc := minimock.NewController(t)

	dockerMock := NewCommonAPIClientMock(mc).ContainerListMock.Set(func(ctx context.Context, options types.ContainerListOptions) (ca1 []types.Container, err error) {
		return []types.Container{mockContainer}, nil
	})

	gw := NewGateway(dockerMock)

	container, err := gw.ContainerGet(entity.ContainerID(mockContainer.ID))

	assert.Nilf(t, err, "ContainerGet returns no error")
	assert.NotNilf(t, container, "Should return a container")

	assert.Equal(t, string(container.ID), mockContainer.ID)
	assert.Equal(t, container.Created, time.Unix(mockContainer.Created, 0),
		"Container creation times match")
	assert.Equal(t, container.Name, strings.TrimLeft(mockContainer.Names[0], "/"),
		"Container names match")
}

func TestGateway_hydrateNetworkFromTypeNetworkSettings(t *testing.T) {
	networkSettings := getMockContainers()[0].NetworkSettings

	networks := hydrateNetworkFromTypeNetworkSettings(*networkSettings)

	i := 0
	for _, nw := range networkSettings.Networks {
		assert.Equal(t, networks[i].ID, nw.NetworkID, "Network ID matches")
		assert.Equal(t, networks[i].Gateway, nw.Gateway, "Gateway matches")
		assert.Equal(t, networks[i].IPAddress, nw.IPAddress, "IP matches")

		i++
	}
}

func TestGateway_hydrateVolumesFromTypeMountPoint(t *testing.T) {
	mounts := getMockContainers()[0].Mounts

	volumes := hydrateVolumesFromTypeMountPoint(mounts)

	for i, m := range mounts {
		assert.Equal(t, volumes[i].Mode, m.Mode, "Mode matches")
		assert.EqualValues(t, volumes[i].Type, m.Type, "Type matches")
		assert.Equal(t, volumes[i].Name, m.Name, "Name matches")
		assert.Equal(t, volumes[i].Destination, m.Destination, "Destination matches")
		assert.Equal(t, volumes[i].Source, m.Source, "Source matches")
		assert.Equal(t, volumes[i].ReadWrite, m.RW, "Read/Write matches")
	}
}

func TestGateway_hydratePortsFromTypePort(t *testing.T) {
	cPorts := getMockContainers()[0].Ports

	ports := hydratePortsFromTypePort(cPorts)

	for i, m := range cPorts {
		assert.Equal(t, ports[i].IP, m.IP, "IP matches")
		assert.Equal(t, ports[i].ContainerPort, m.PrivatePort, "Container Port matches")
		assert.Equal(t, ports[i].HostPort, m.PublicPort, "Host Port matches")
		assert.Equal(t, ports[i].Protocol, m.Type, "Protocol matches")
	}
}

func TestGateway_hydrateImageFromTypeContainerWithNoTag(t *testing.T) {
	mockContainer := &types.Container{
		Image:   "foo/bar",
		ImageID: "sha256:123456",
	}

	image := hydrateImageFromTypeContainer(mockContainer)

	assert.Equal(t, image.Name, "foo/bar", "Image name matches")
	assert.Equal(t, image.Tag, "latest", "Image tag matches")
	assert.Equal(t, image.ID, mockContainer.ImageID, "Image ID matches")
}

func TestGateway_hydrateImageFromTypeContainerWithDefinedTag(t *testing.T) {
	mockContainer := &types.Container{
		Image:   "foo/bar:1.2.3",
		ImageID: "sha256:123456",
	}

	image := hydrateImageFromTypeContainer(mockContainer)

	assert.Equal(t, image.Name, "foo/bar", "Image name matches")
	assert.Equal(t, image.Tag, "1.2.3", "Image tag matches")
	assert.Equal(t, image.ID, mockContainer.ImageID, "Image ID matches")
}

func TestGateway_ContainerGetWithEmptyContainers(t *testing.T) {
	mc := minimock.NewController(t)

	dockerMock := NewCommonAPIClientMock(mc).ContainerListMock.Set(func(ctx context.Context, options types.ContainerListOptions) (ca1 []types.Container, err error) {
		return []types.Container{}, nil
	})

	gw := NewGateway(dockerMock)

	container, err := gw.ContainerGet(entity.ContainerID("abcd"))

	assert.Nilf(t, err, "Should not return errors")
	assert.Nilf(t, container, "Should return nil if no containers were found")
}
