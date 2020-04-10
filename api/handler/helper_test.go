package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddContentType(t *testing.T) {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.HandlerFunc(AddContentType),
	)

	r.Handle("/", n)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestGetNotFoundHandler(t *testing.T) {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(GetNotFoundHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/404")
	assert.Nilf(t, err, "Response should not return an error")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nilf(t, err, "Body reading should not return an error")

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.JSONEq(t, string(body), "{\"error\": \"Not found\"}")
}
