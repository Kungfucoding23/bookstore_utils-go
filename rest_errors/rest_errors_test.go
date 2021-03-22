package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("internal error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "internal_server_error", err.Error())
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "internal error", err.Causes()[0])
}
func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}
func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "not_found", err.Error())
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("this is the message", http.StatusNotImplemented, "not_implemented", nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "not_implemented", err.Error())
	assert.Nil(t, err.Causes())
}
func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "unauthorized", err.Error())
}

func TestNewRestErrorFromBytesValidJSON(t *testing.T) {
	var jsonStr = []byte(`{
		"id"  : 15,
		"foo" : { "foo": 123, "bar": "baz" }
	}`)
	_, err := NewRestErrorFromBytes(jsonStr)
	assert.Nil(t, err)
}
func TestNewRestErrorFromBytesInvalidJSON(t *testing.T) {
	var jsonStr = []byte(``)
	_, err := NewRestErrorFromBytes(jsonStr)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid json", err.Error())
}
