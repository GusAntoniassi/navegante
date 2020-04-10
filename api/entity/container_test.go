package entity

import (
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// @TODO: Put this function somewhere more generic, to allow entity and handler tests to use it
func makeMockContainers() []*entity.Container {
	return []*entity.Container{
		{
			ID:         "0123abcd456e",
			Cmd:        []string{"echo", "foo"},
			Entrypoint: []string{"/bin/bash"},
			Created:    time.Date(2001, 01, 01, 01, 01, 01, 01, time.UTC),
			Name:       "smart_einstein",
			State:      "running",
			Status:     "Up 10 hours",
			Image: &entity.Image{
				ID:   "sha256:abc123",
				Name: "registry.foo.bar/foo",
				Tag:  "latest",
			},
			Ports: []entity.PortMapping{
				{
					IP:            "0.0.0.0",
					ContainerPort: 80,
					HostPort:      8080,
					Protocol:      "TCP",
				},
			},
			Labels: nil,
			Volumes: []entity.Volume{
				{
					Name:        "foobar",
					Type:        "volume",
					Source:      "/var/lib/docker/volumes/foobar/_data",
					Destination: "/var/lib/foo",
					Mode:        "rw",
					ReadWrite:   true,
				},
			},
			Networks: []entity.Network{
				{
					ID:        "1234",
					Name:      "foobar",
					Gateway:   "172.0.0.1",
					IPAddress: "172.0.0.2",
					Links:     nil,
					Aliases:   nil,
				},
			},
		},
		{
			ID:         "0321dcba654e",
			Cmd:        []string{"echo", "bar"},
			Entrypoint: []string{"/bin/bash"},
			Created:    time.Date(2001, 01, 01, 01, 01, 01, 01, time.UTC),
			Name:       "silly_bach",
			State:      "running",
			Status:     "Up 4 hours",
			Image: &entity.Image{
				ID:   "sha256:abcde123123",
				Name: "registry.foo.bar/bar",
				Tag:  "1.0.2",
			},
			Ports: []entity.PortMapping{
				{
					ContainerPort: 80,
					Protocol:      "TCP",
				},
			},
		},
	}
}

func assertContainersAreEqual(t *testing.T, c *entity.Container, apiC Container) {
	ports := make([]string, len(c.Ports))
	volumes := make([]string, len(c.Volumes))
	networks := make([]string, len(c.Networks))

	for i, p := range c.Ports {
		ports[i] = p.String()
	}

	for i, v := range c.Volumes {
		volumes[i] = v.String()
	}

	for i, n := range c.Networks {
		networks[i] = n.String()
	}

	assert.Equal(t, c.ID, apiC.ID)
	assert.Equal(t, c.Name, apiC.Name)
	assert.Equal(t, c.Image.String(), apiC.Image)
	assert.Equal(t, c.Cmd.String(), apiC.Cmd)
	assert.Equal(t, c.Entrypoint.String(), apiC.Entrypoint)
	assert.Equal(t, volumes, apiC.Volumes)
	assert.Equal(t, ports, apiC.Ports)
	assert.Equal(t, networks, apiC.Networks)
}

func TestNewContainer(t *testing.T) {
	mc := makeMockContainers()

	for _, c := range mc {
		apiC := NewContainer(c)
		assertContainersAreEqual(t, c, apiC)
	}
}

func TestNewContainerWithEmptyValues(t *testing.T) {
	c := &entity.Container{
		ID:         "0321dcba654e",
		Cmd:        []string{},
		Entrypoint: []string{"/bin/bash"},
		Created:    time.Date(2001, 01, 01, 01, 01, 01, 01, time.UTC),
		Name:       "silly_bach",
		State:      "running",
		Status:     "Up 4 hours",
		Image: &entity.Image{
			ID:   "sha256:abcde123123",
			Name: "registry.foo.bar/bar",
			Tag:  "1.0.2",
		},
	}

	apiC := NewContainer(c)
	assert.Equal(t, apiC.Cmd, "")
	assert.Empty(t, apiC.Networks)
	assert.Empty(t, apiC.Ports)
	assert.Empty(t, apiC.Volumes)
}
