package handler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_formatJsonError(t *testing.T) {
	message := "foobar"
	expected := fmt.Sprintf("{\"message\": \"%s\"}", message)

	json := string(formatJSONError(message))

	assert.JSONEq(t, expected, json)
}
