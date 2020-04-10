package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gojuno/minimock/v3"
	"github.com/gorilla/mux"
	apiEntity "github.com/gusantoniassi/navegante/api/entity"
	"github.com/gusantoniassi/navegante/core/entity"
	"github.com/gusantoniassi/navegante/gateway/containergateway"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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

func getTestServer(gw containergateway.Gateway) *httptest.Server {
	r := mux.NewRouter()
	n := negroni.New()

	MakeContainerHandlers(r, n, gw)

	ts := httptest.NewServer(r)

	return ts
}

func TestContainer_getAllContainers(t *testing.T) {
	mc := minimock.NewController(t)
	mockContainers := makeMockContainers()
	gw := NewContainerMock(mc).ContainerGetAllMock.Set(func() ([]*entity.Container, error) {
		return mockContainers, nil
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers")
	assert.Nilf(t, err, "http.Get should not return an error")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nilf(t, err, "Body reading should not return an error")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, body)

	var containers []apiEntity.Container
	err = json.Unmarshal(body, &containers)
	assert.Nilf(t, err, "Response JSON unmarshalling should not return errors")

	for i, c := range containers {
		assert.Equal(t, mockContainers[i].ID, c.ID)
	}
}

func TestContainer_getAllContainersEmpty(t *testing.T) {
	mc := minimock.NewController(t)

	gw := NewContainerMock(mc).ContainerGetAllMock.Set(func() ([]*entity.Container, error) {
		return []*entity.Container{}, nil
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers")
	assert.Nilf(t, err, "http.Get should not return an error")

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.NotEmpty(t, resp.Body)
}

func TestContainer_getAllContainersError(t *testing.T) {
	mc := minimock.NewController(t)

	gw := NewContainerMock(mc).ContainerGetAllMock.Set(func() ([]*entity.Container, error) {
		return nil, errors.Errorf("foobar")
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers")
	assert.Nilf(t, err, "http.Get should not return an error")

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.NotEmpty(t, resp.Body)
}

func TestContainer_getContainer(t *testing.T) {
	mc := minimock.NewController(t)
	mockContainers := makeMockContainers()
	gw := NewContainerMock(mc).ContainerGetMock.Set(func(cid entity.ContainerID) (*entity.Container, error) {
		return mockContainers[0], nil
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers/0123abcd456e")
	assert.Nilf(t, err, "http.Get should not return an error")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nilf(t, err, "Body reading should not return an error")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, body)

	var container apiEntity.Container
	err = json.Unmarshal(body, &container)
	assert.Nilf(t, err, "Response JSON unmarshalling should not return errors")

	assert.Equal(t, mockContainers[0].ID, container.ID)
}

func TestContainer_getContainerNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	gw := NewContainerMock(mc).ContainerGetMock.Set(func(cid entity.ContainerID) (*entity.Container, error) {
		return nil, nil
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers/0123abcd456e")
	assert.Nilf(t, err, "http.Get should not return an error")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nilf(t, err, "Body reading should not return an error")

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.NotEmpty(t, body)
}

func TestContainer_getContainerError(t *testing.T) {
	mc := minimock.NewController(t)
	gw := NewContainerMock(mc).ContainerGetMock.Set(func(cid entity.ContainerID) (*entity.Container, error) {
		return nil, errors.Errorf("foobar")
	})

	ts := getTestServer(gw)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/containers/0123abcd456e")
	assert.Nilf(t, err, "http.Get should not return an error")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nilf(t, err, "Body reading should not return an error")

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.NotEmpty(t, body)
}
