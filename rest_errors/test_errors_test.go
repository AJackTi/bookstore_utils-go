package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := New(http.StatusInternalServerError, errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "The server encountered an unexpected condition that prevented it from fulfilling the request.", err.Detail)
	assert.EqualValues(t, "INTERNAL SERVER ERROR", err.Title)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0].(error).Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := New(http.StatusBadRequest, errors.New("invalid argument"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "The server cannot or will not process the request due to something that is perceived to be a client error", err.Detail)
	assert.EqualValues(t, "BAD REQUEST", err.Title)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "invalid argument", err.Causes[0].(error).Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := New(http.StatusNotFound, errors.New("not found"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "The origin server did not find a current representation for the target resource or is not willing to disclose that one exists.", err.Detail)
	assert.EqualValues(t, "NOT FOUND", err.Title)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "not found", err.Causes[0].(error).Error())
}

func TestNewError(t *testing.T) {
	err := New(http.StatusBadRequest, errors.New("invalid database"), "Something went wrong")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "Something went wrong", err.Detail)
	assert.EqualValues(t, "Something went wrong", err.Title)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "invalid database", err.Causes[0].(error).Error())
}
