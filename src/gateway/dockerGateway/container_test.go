package dockerGateway

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/gojuno/minimock/v3"
	"github.com/gusantoniassi/shipmate/core/entity"
	"github.com/stretchr/testify/assert"
)

func getMockContainers() []types.Container {
	return []types.Container{
		{
			ID:         "0123abcd456e",
			Names:      []string{"/smart_einstein"},
			Image:      "registry.foo.bar/foo:latest",
			ImageID:    "sha256:abcde123456",
			Command:    "/bin/bash echo foo",
			Created:    1583610097,
			Ports:      []types.Port{},
			SizeRw:     0,
			SizeRootFs: 0,
			Labels:     nil,
			State:      "running",
			Status:     "Up 4 hours",
			HostConfig: struct {
				NetworkMode string `json:",omitempty"`
			}{},
			NetworkSettings: nil,
			Mounts:          nil,
		},
		{
			ID:      "0321dcba654e",
			Names:   []string{"/silly_bach"},
			Image:   "registry.foo.bar/bar:latest",
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
		assert.Equal(t, string(container.Id), mockContainers[i].ID,
			"Container IDs match")
		assert.Equal(t, container.Created, time.Unix(mockContainers[i].Created, 0),
			"Container creation times match")
		assert.Equal(t, container.Name, strings.TrimLeft(mockContainers[i].Names[0], "/"),
			"Container names match")
	}
}
